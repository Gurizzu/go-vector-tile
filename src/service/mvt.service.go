package service

import (
	"context"
	"fmt"
	"log"
	"vector-tile/src/config"
	"vector-tile/src/utils"

	"github.com/paulmach/orb/encoding/mvt"
)

//* Change o service content with ALT+C(Case sensitive) then CTRL+D o:
//? role and Mvt

type MvtService struct {
	ctx       context.Context
	upostgres *config.PostgresDbUtil
}

func NewMvtService() *MvtService {
	o := &MvtService{
		ctx:       context.Background(),
		upostgres: config.NewPostgresDbUtil(),
	}

	return o
}

func (s *MvtService) Get(x, y, z int) (data []byte, errMessage string) {

	envelope := utils.GenerateEnvelope(x, y, z)

	bound := fmt.Sprintf("ST_Segmentize(ST_MakeEnvelope(%f, %f, %f, %f, 3857), %f)",
		envelope["x_min"], envelope["y_min"], envelope["x_max"], envelope["y_max"],
		(envelope["x_max"]-envelope["x_min"])/4)

	db, err := s.upostgres.Connect()
	if err != nil {
		log.Println(err)
		data = nil
		errMessage = err.Error()
		return
	}

	query := fmt.Sprintf(`
                WITH 
                    bounds AS (
                        SELECT 
                            %s AS geom,
                            %s::box2d AS b2d
                    ),
                    datas AS (
                        SELECT
                            *
                        FROM
                            indonesia_prov
                    ),
                    mvtgeom AS (
                        SELECT 
                            ST_AsMVTGeom(ST_Transform(datas.geometry , 3857) , bounds.b2d) AS geom, 
                            datas.*
                        FROM
                            bounds,
                            datas
                        WHERE
                            datas.geometry && ST_Transform(bounds.geom, 4326)
                    )
                SELECT 
                    ST_AsMVT(mvtgeom.*, 'geojsonLayer')
                FROM 
                    mvtgeom
            `, bound, bound)

	var mvtBytes []interface{}

	err = db.Raw(query).Scan(&mvtBytes).Error
	if err != nil {
		log.Println(err)
		data = nil
		errMessage = err.Error()
		return
	}

	layer, err := mvt.Unmarshal(mvtBytes[0].([]byte))
	if err != nil {
		log.Println(err)
		return
	}

	layerByte, err := mvt.Marshal(layer)
	if err != nil {
		log.Println(err)
		data = nil
		errMessage = err.Error()
		return
	}

	return layerByte, ""
}
