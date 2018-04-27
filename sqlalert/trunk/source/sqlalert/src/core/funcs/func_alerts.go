// Functions for alert output.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package funcs

import (
	"core/script"
	"core/tools"
	"core/value"
	"errors"
	"fmt"
)

// alertElastic() - Aggregate the values to a list.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func alertElastic(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d(expected 1)", len(args))
		return nil, errors.New(msg)
	}

	list, ok := value.AsList(args[0])
	if !ok || len(list) == 0 {
		return nil, nil
	}

	list = value.CopyList(list)
	for cc, item := range list {
		dict, ok := value.AsDict(item)
		if !ok {
			msg := fmt.Sprintf("list[%d] not a dict", cc)
			return nil, errors.New(msg)
		}

		if val := ctx.GetX("__level__"); val != nil {
			if val, err := levelName([]interface{}{val}, ctx); err == nil {
				if val != nil {
					dict["level"] = val
				}
			}
		}

		if val := ctx.GetX("__type__"); val != nil {
			dict["type"] = val
		}

		if val := ctx.GetX("__subtype__"); val != nil {
			dict["subtype"] = val
		}

		if val := ctx.GetX("__alert_es_withdesc__"); value.IsTrue(val) {
			if val, ok := ctx.GetXStr("__desc_type__"); ok && val != "" {
				dict["desc_type"] = val
			}

			if val, ok := ctx.GetXStr("__desc_subtype__"); ok && val != "" {
				dict["desc_subtype"] = val
			}
		}
	}

	return elasticInsert([]interface{}{list}, ctx)
}

func __alertEmailDesc(ctx *script.Cntx) interface{} {
	listDesc, ok := ctx.GetXList("__email_desc_list__")
	if !ok {
		return []interface{}{}
	}

	listRes := []interface{}{}
	for _, item := range listDesc {
		if dict, ok := value.AsDict(item); ok {
			if key, ok := dict["name"]; !ok {
				listRes = append(listRes, map[string]interface{}{})
			} else {
				if val, ok := dict["value"]; ok {
					if ok, err := miscIsFunc([]interface{}{val}, ctx); value.IsTrue(ok) && err == nil {
						if val, err = miscCall([]interface{}{val}, ctx); err == nil {
							listRes = append(listRes, map[string]interface{}{"name": key, "value": val})
						}
					} else {
						if varVal := ctx.Get(value.ToStr(val)); varVal != nil {
							val = varVal
						}

						listRes = append(listRes, map[string]interface{}{"name": key, "value": val})
					}
				} else {
					listRes = append(listRes, map[string]interface{}{"name": key})
				}
			}
		}
	}

	return listRes
}

func __alertEmailFieldDesc(item interface{}, ctx *script.Cntx) interface{} {
	listKeys, ok := ctx.GetXList("__fields__")
	if !ok || len(listKeys) == 0 {
		if dict, ok := value.AsDict(item); ok {
			if res, err := miscKeys([]interface{}{dict}, ctx); err == nil {
				listKeys = value.List(res)
			}
		}
	}

	fieldAll, listRes := ctx.GetX("__fields_all__"), []interface{}{}
	if dict, ok := value.AsDict(fieldAll); !ok {
		for _, name := range listKeys {
			listRes = append(listRes, map[string]interface{}{"name": name, "desc": name})
		}
	} else {
		for _, name := range listKeys {
			key := value.ToStr(name)

			if item, ok := dict[key]; ok && value.IsDict(item) {
				dictCpy := value.CopyDict(item)
				dictCpy["name"] = name
				listRes = append(listRes, dictCpy)
			} else {
				listRes = append(listRes, map[string]interface{}{"name": name, "desc": name})
			}
		}
	}

	return listRes
}

func alertEmail(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d(expected 1)", len(args))
		return nil, errors.New(msg)
	}

	levelItem := __levelItem(nil, ctx)
	if dict, ok := ctx.GetXDict("__levels_sendemail__"); ok {
		if value.IsFalse(dict[value.ToStr(levelItem["name"])]) {
			return nil, nil
		}
	}

	listAlert := args[0]
	if val := ctx.GetX("__enable_alert_throttle__"); value.IsTrue(val) {
		listAlert, _ = throttleCheck([]interface{}{listAlert}, ctx)
	}

	list := value.CopyList(listAlert)
	if len(list) == 0 {
		return nil, nil
	}

	if val, ok := ctx.GetXStr("__sort_key__"); ok && val != "" {
		list = tools.Sort(list, val, ">=")
	}

	ctx.Set("__email_desc__", __alertEmailDesc(ctx))
	ctx.Set("__email_fields__", __alertEmailFieldDesc(list[0], ctx))

	if val := ctx.GetX("__desc_title__"); val != nil {
		ctx.Set("__email_title__", val)
	}

	return emailSend([]interface{}{list}, ctx)
}

func __alertOwnerMapItem(data map[string]interface{}, key string, list []interface{}, isIP bool, ctx *script.Cntx) map[string]bool {
	dataValue, ok := data[key]
	if !ok {
		return nil
	}

	fields := map[string]bool{}
	for _, item := range list {
		if itemDict, ok := value.AsDict(item); ok {
			if itemValue, ok := itemDict[key]; ok {
				match := false

				if value.Compare(dataValue, "==", itemValue) {
					match = true
				} else if isIP {
					res, err := miscIpIn([]interface{}{dataValue, itemValue}, ctx)
					if err == nil && value.IsTrue(res) {
						match = true
					}
				}

				for k, v := range itemDict {
					if k != key {
						if match {
							data[k] = v
						} else {
							data[k] = ""
						}
						fields[k] = true
					}
				}

				if match {
					break
				}
			}
		}
	}

	return fields
}

func __alertOwnerMap(list []interface{}, key string, ctx *script.Cntx) ([]interface{}, error) {
	data, ok := ctx.GetXList("__" + key + "list__")
	if !ok || data == nil || len(data) == 0 {
		return list, nil
	}

	isIP := false
	if dict, ok := ctx.GetXDict("__fields_ipfields__"); ok {
		if val, ok := dict[key]; ok && value.IsTrue(val) {
			isIP = true
		}
	}

	fields := map[string]bool{}
	for _, item := range list {
		dict, ok := value.AsDict(item)
		if !ok {
			continue
		}

		subFields := __alertOwnerMapItem(dict, key, data, isIP, ctx)
		for k, _ := range subFields {
			fields[k] = true
		}
	}

	if alertFields, ok := ctx.GetXList("__fields__"); ok && len(fields) != 0 {
		for k, _ := range fields {
			alertFields = append(alertFields, k)
		}

		ctx.Set("__fields__", alertFields)
	}

	return list, nil
}

func alertAlert(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d(expected 1)", len(args))
		return nil, errors.New(msg)
	}

	list, err := args[0], error(nil)
	if val := ctx.GetX("__enable_alert_hisdata__"); value.IsTrue(val) {
		if list, err = hisCheck([]interface{}{list}, ctx); err != nil {
			return nil, err
		}
	}

	newlist := value.Copy(list)
	if val := ctx.GetX("__enable_owner_sip__"); value.IsTrue(val) {
		if newlist, err = __alertOwnerMap(value.List(newlist), "sip", ctx); err != nil {
			return nil, err
		}
	}

	if val := ctx.GetX("__enable_alert_es__"); value.IsTrue(val) {
		alertElastic([]interface{}{newlist}, ctx)
	}

	if val := ctx.GetX("__enable_alert_email__"); value.IsTrue(val) {
		alertEmail([]interface{}{newlist}, ctx)
	}

	if val := ctx.GetX("__enable_alert_debug__"); value.IsTrue(val) {
		miscPrintList([]interface{}{newlist}, ctx)
	}

	return list, nil
}
