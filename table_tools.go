package arango_tools

import (
	"github.com/apex/log"
	"github.com/thedanielforum/aranGO"
)

// CreateCollections creates collections if they don't exist
func CreateTables(arango *aranGO.Session, db string, tables ...string) {
	for _, table := range tables {
		if !arango.DB(db).ColExist(table) {
			doc := aranGO.NewCollectionOptions(table, true)
			arango.DB(db).CreateCollection(doc)
			log.WithFields(log.Fields{
				"table": table,
			}).Info("creating collection")
		} else {
			log.WithFields(log.Fields{
				"table": table,
			}).Info("collection already exists")
		}
	}
}

// CreateEdges creates edges if they don't exist
func CreateEdges(arango *aranGO.Session, db string, edges ...string) {
	for _, edge := range edges {
		if !arango.DB(db).ColExist(edge) {
			doc := aranGO.NewCollectionOptions(edge, true)
			doc.IsEdge()
			log.WithFields(log.Fields{
				"table": edge,
			}).Info("creating edge")
		} else {
			log.WithFields(log.Fields{
				"table": edge,
			}).Info("edge already exists")
		}
	}
}
