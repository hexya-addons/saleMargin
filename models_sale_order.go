package sale_margin

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

//import odoo.addons.decimal_precision as dp
func init() {
	h.SaleOrderLine().DeclareModel()

	h.SaleOrderLine().AddFields(map[string]models.FieldDefinition{
		"Margin": models.FloatField{
			Compute: h.SaleOrderLine().Methods().ProductMargin(),
			//digits=dp.get_precision('Product Price')
			Stored: true,
		},
		"PurchasePrice": models.FloatField{
			String: "Cost",
			//digits=dp.get_precision('Product Price')
		},
	})
	h.SaleOrderLine().Methods().ComputeMargin().DeclareMethod(
		`ComputeMargin`,
		func(rs m.SaleOrderLineSet, order_id interface{}, product_id interface{}, product_uom_id interface{}) {
			//        frm_cur = self.env.user.company_id.currency_id
			//        to_cur = order_id.pricelist_id.currency_id
			//        purchase_price = product_id.standard_price
			//        if product_uom_id != product_id.uom_id:
			//            purchase_price = product_id.uom_id._compute_price(
			//                purchase_price, product_uom_id)
			//        ctx = self.env.context.copy()
			//        ctx['date'] = order_id.date_order
			//        price = frm_cur.with_context(ctx).compute(
			//            purchase_price, to_cur, round=False)
			//        return price
		})
	h.SaleOrderLine().Methods().GetPurchasePrice().DeclareMethod(
		`GetPurchasePrice`,
		func(rs m.SaleOrderLineSet, pricelist interface{}, product interface{}, product_uom interface{}, date interface{}) {
			//        frm_cur = self.env.user.company_id.currency_id
			//        to_cur = pricelist.currency_id
			//        purchase_price = product.standard_price
			//        if product_uom != product.uom_id:
			//            purchase_price = product.uom_id._compute_price(
			//                purchase_price, product_uom)
			//        ctx = self.env.context.copy()
			//        ctx['date'] = date
			//        price = frm_cur.with_context(ctx).compute(
			//            purchase_price, to_cur, round=False)
			//        return {'purchase_price': price}
		})
	h.SaleOrderLine().Methods().ProductIdChangeMargin().DeclareMethod(
		`ProductIdChangeMargin`,
		func(rs m.SaleOrderLineSet) {
			//        if not self.order_id.pricelist_id or not self.product_id or not self.product_uom:
			//            return
			//        self.purchase_price = self._compute_margin(
			//            self.order_id, self.product_id, self.product_uom)
		})
	h.SaleOrderLine().Methods().Create().Extend(
		`Create`,
		func(rs m.SaleOrderLineSet, vals models.RecordData) {
			//        vals.update(self._prepare_add_missing_fields(vals))
			//        if 'purchase_price' not in vals:
			//            order_id = self.env['sale.order'].browse(vals['order_id'])
			//            product_id = self.env['product.product'].browse(vals['product_id'])
			//            product_uom_id = self.env['product.uom'].browse(
			//                vals['product_uom'])
			//
			//            vals['purchase_price'] = self._compute_margin(
			//                order_id, product_id, product_uom_id)
			//        return super(SaleOrderLine, self).create(vals)
		})
	h.SaleOrderLine().Methods().ProductMargin().DeclareMethod(
		`ProductMargin`,
		func(rs h.SaleOrderLineSet) h.SaleOrderLineData {
			//        for line in self:
			//            currency = line.order_id.pricelist_id.currency_id
			//            price = line.purchase_price
			//            line.margin = currency.round(
			//                line.price_subtotal - (price * line.product_uom_qty))
		})
	h.SaleOrder().DeclareModel()

	h.SaleOrder().AddFields(map[string]models.FieldDefinition{
		"Margin": models.MonetaryField{
			Compute: h.SaleOrder().Methods().ProductMargin(),
			Help: "It gives profitability by calculating the difference between" +
				"the Unit Price and the cost.",
			//currency_field='currency_id'
			//digits=dp.get_precision('Product Price')
			Stored: true,
		},
	})
	h.SaleOrder().Methods().ProductMargin().DeclareMethod(
		`ProductMargin`,
		func(rs h.SaleOrderSet) h.SaleOrderData {
			//        for order in self:
			//            order.margin = sum(order.order_line.filtered(
			//                lambda r: r.state != 'cancel').mapped('margin'))
		})
}
