package sale_margin

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.SaleReport().DeclareModel()

	h.SaleReport().AddFields(map[string]models.FieldDefinition{
		"Margin": models.FloatField{
			String: "Margin",
		},
	})
	h.SaleReport().Methods().Select().DeclareMethod(
		`Select`,
		func(rs m.SaleReportSet) {
			//        return super(SaleReport, self)._select() + ", SUM(l.margin / COALESCE(cr.rate, 1.0)) AS margin"
		})
}
