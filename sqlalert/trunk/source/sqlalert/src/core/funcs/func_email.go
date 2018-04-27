// Functions for email.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package funcs

import (
	"bytes"
	"core/email"
	"core/script"
	"core/value"
	"errors"
	"fmt"
)

// emailCheckSmtp() - Check the __email_smtp__ configuration.
//
// @ctx: Script context.
//
// This functino returns a DICT if __email_smtp__ is valid, otherwise
// it returns nil and error info.
var g_email_smtp_strkeys = []string{"host", "address", "password"}

func emailCheckSmtp(ctx *script.Cntx) (map[string]interface{}, error) {
	dict, ok := ctx.GetX("__email_smtp__").(map[string]interface{})
	if !ok || len(dict) == 0 {
		msg := "'__email_smtp__' not found or invalid DICT"
		return nil, errors.New(msg)
	}

	for _, item := range g_email_smtp_strkeys {
		if itemValue, ok := dict[item]; !ok || !value.IsStr(itemValue) || !value.IsTrue(itemValue) {
			msg := fmt.Sprintf("'__email_smtp__['%s']' not found or empty", value.ToStr(item))
			return nil, errors.New(msg)
		}
	}

	if port, ok := dict["port"].(int64); !ok || value.IsFalse(port) {
		return nil, errors.New(fmt.Sprintf("smtp['port'] not found or empty"))
	}

	return dict, nil
}

// emailCheckSendto() - Check the __email_sendto__ configuration.
//
// @ctx: Script context.
//
// This function check __email_sendto__ configuration and returns a list of string.
func emailCheckSendto(ctx *script.Cntx) ([]string, error) {
	ctxValue := ctx.GetX("__email_sendto__")
	if value.IsFalse(ctxValue) {
		msg := fmt.Sprintf("'__email_sendto__' is empty: %s", value.ToStr(ctxValue))
		return nil, errors.New(msg)
	}

	switch ctxValue.(type) {
	case string:
		return []string{ctxValue.(string)}, nil

	case []interface{}:
		list := []string{}
		for _, item := range ctxValue.([]interface{}) {
			if !value.IsStr(item) || value.IsFalse(item) {
				msg := fmt.Sprintf("'__email_sendto__'[%d] not STRING or empty", value.ToStr(item))
				return nil, errors.New(msg)
			}

			list = append(list, item.(string))
		}
		return list, nil
	}

	msg := fmt.Sprintf("'__email_sendto__' not STRING or list of STRING: %s", value.ToStr(ctxValue))
	return nil, errors.New(msg)
}

// checkTitle() - Check the __email_title__ configuration.
//
// @ctx: Script context.
//
// This function check __email_title__ configuration and returns a string.
var g_email_default_title = "You have an alert message to process"

func emailCheckTitle(ctx *script.Cntx) string {
	title, ok := ctx.GetXStr("__email_title__")
	if !ok || title == "" {
		return g_email_default_title
	}

	return title
}

// emailCheckDesc() - Check the __email_desc__ configuration.
//
// @ctx: Script context.
//
// This function check __email_desc__ configuration and returns
// a list of DICT.
func emailCheckDesc(ctx *script.Cntx) []map[string]interface{} {
	listValue, ok := ctx.GetXList("__email_desc__")
	if !ok || len(listValue) == 0 {
		return nil
	}

	list := []map[string]interface{}{}
	for _, item := range listValue {
		if value.IsDict(item) {
			list = append(list, item.(map[string]interface{}))
		}
	}

	return list
}

// emailCheckFields() - Check the __email_fields__ configuration.
//
// @ctx: Script context.
//
// This function check __email_fields__ configuration and returns
// a list of DICT.
func emailCheckFields(ctx *script.Cntx) []map[string]interface{} {
	listValue, ok := ctx.GetXList("__email_fields__")
	if !ok || len(listValue) == 0 {
		return nil
	}

	list := []map[string]interface{}{}
	for _, item := range listValue {
		if value.IsDict(item) {
			list = append(list, item.(map[string]interface{}))
		}
	}

	return list
}

// emailFormatMsg() - Format email message.
//
// @buffer: Bytes buffer.
// @desc:   List of email descrition fields.
// @ctx:  Script context.
func emailFormatDesc(buffer *bytes.Buffer, desc []map[string]interface{}, ctx *script.Cntx) {
	if len(desc) == 0 {
		return
	}

	buffer.WriteString("<br/>")
	buffer.WriteString("<table border=\"0\" cellspacing=\"0px\" cellpadding=\"2px\" style=\"text-align:left; border-collapse:collapse\">")
	for _, item := range desc {
		buffer.WriteString("<tr/>")

		if value.IsTrue(item) {
			descName, descValue := "", ""
			if v, ok := item["name"]; ok {
				descName = value.ToStr(v)
			}
			if v, ok := item["value"]; ok {
				descValue = value.ToStr(v)
			}

			buffer.WriteString(fmt.Sprintf("<td><b>%s:</b></td><td>%s</td>", descName, descValue))
		} else {
			buffer.WriteString("<td>&nbsp;</td><td>&nbsp;</td>")
		}

		buffer.WriteString("<tr>")
	}
	buffer.WriteString("</table>")
}

var g_format_funcs = map[string]struct {
	name string
	arg  interface{}
}{
	"time_ms":    {"fmt_time", "ms"},
	"time_us":    {"fmt_time", "us"},
	"time_s":     {"fmt_time", "s"},
	"bytes":      {"fmt_bytes", nil},
	"bits":       {"fmt_bits", nil},
	"percentage": {"fmt_percentage", nil},
	"pct":        {"fmt_percentage", nil},
	"int":        {"fmt_int", nil},
	"float":      {"fmt_float", 2},
	"str":        {"fmt_str", 40},
}

// emailFormatField() - Format fiead to readable string.
//
// @buffer: Bytes buffer.
// @fields: List of data format fileds.
// @data:   List of DICT to format.
// @ctx:    Script context.
func emailFormatField(v, fmtValue interface{}, ctx *script.Cntx) string {
	fmtStr, fmtArgs := value.ToStr(fmtValue), []interface{}{v}

	if item, ok := g_format_funcs[fmtStr]; ok {
		fmtStr = item.name
		if item.arg != nil {
			fmtArgs = append(fmtArgs, item.arg)
		}
	}

	if token, ok := ctx.Defines[fmtStr]; ok {
		ctx.PushAndRefer()
		result, err := script.ExecDefineFunc(token, fmtArgs, ctx)
		ctx.PopAndDefer()

		if err == nil {
			return value.ToStr(result)
		}
	} else if function, ok := ctx.Funcs[fmtStr]; ok {
		if result, err := function(fmtArgs, ctx); err == nil {
			return value.ToStr(result)
		}
	}

	return value.ToStr(v)
}

// emailFormatMsg() - Format email message.
//
// @buffer: Bytes buffer.
// @fields: List of data format fileds.
// @data:   List of DICT to format.
// @ctx:    Script context.
func emailFormatData(buffer *bytes.Buffer, fields, data []map[string]interface{}, ctx *script.Cntx) {
	buffer.WriteString("<table border=\"1\" cellspacing=\"0px\" cellpadding=\"5px\" style=\"text-align:left; border-collapse:collapse\">")
	buffer.WriteString("<tr>")
	buffer.WriteString("<td></td>")

	fieldList := []map[string]interface{}{}
	for _, item := range fields {
		if _, ok := item["name"]; !ok {
			continue
		}

		fieldList = append(fieldList, item)
		name := value.ToStr(item["name"])
		if desc, ok := item["desc"]; ok {
			buffer.WriteString("<td style=\"text-align: center\"><b>")
			buffer.WriteString(value.ToStr(desc))
			buffer.WriteString("</b>")
		} else {
			buffer.WriteString("<td><b>")
			buffer.WriteString(name)
			buffer.WriteString("</b></td>")
		}
	}
	buffer.WriteString("</tr>")

	for cc, dataItem := range data {
		buffer.WriteString("<tr>")
		buffer.WriteString("<td style=\"text-align:right\">")
		buffer.WriteString(fmt.Sprintf("%d", cc+1))
		buffer.WriteString("</td>")

		for _, item := range fieldList {
			if _, ok := item["name"]; !ok {
				continue
			}

			name := value.ToStr(item["name"])
			if nameValue, ok := dataItem[name]; ok {
				buffer.WriteString("<td>")
				if format, ok := item["fmt"]; ok {
					str := emailFormatField(nameValue, format, ctx)
					buffer.WriteString(str)
				} else {
					buffer.WriteString(value.ToStr(nameValue))
				}
				buffer.WriteString("</td>")
			} else {
				buffer.WriteString("<td></td>")
			}
		}
		buffer.WriteString("</tr>")
	}

	buffer.WriteString("</table>")
}

// emailFormatFooter() - Format email footer message.
//
// @buffer: Bytes buffer.
// @ctx:    Script context.
func emailFormatFooter(buffer *bytes.Buffer, ctx *script.Cntx) {
	if msg, ok := ctx.GetXStr("__email_footer__"); ok {
		buffer.WriteString(msg)
	}
}

// emailFormatMsg() - Format email message.
//
// @from: Address of sender.
// @data: Data to send.
// @ctx:  Script context.
func emailFormatMessage(from string, data []map[string]interface{}, ctx *script.Cntx) (*email.Message, error) {
	sendto, err := emailCheckSendto(ctx)
	if err != nil {
		return nil, err
	}

	title := emailCheckTitle(ctx)
	desc := emailCheckDesc(ctx)
	fields := emailCheckFields(ctx)

	msg := email.NewMessage()
	msg.SetHeader("To", sendto...)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", "[SQLAlert] "+title)

	buffer := &bytes.Buffer{}
	buffer.WriteString("<html>")
	buffer.WriteString("<body>")

	emailFormatDesc(buffer, desc, ctx)
	emailFormatData(buffer, fields, data, ctx)
	emailFormatFooter(buffer, ctx)

	buffer.WriteString("</body>")
	buffer.WriteString("</html>")

	msg.SetBody("text/html", buffer.String())
	return msg, nil
}

func __emailSend(host string, port int, addr, pwd string, msg *email.Message) {
	email.NewPlainDialer(host, port, addr, pwd).DialAndSend(msg)
}

// emailSend() - Send email with given data.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func emailSend(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	extData, ok := ctx.GetXDict("__alert_append__")
	if !ok {
		extData = nil
	}

	data := []map[string]interface{}{}
	switch args[0].(type) {
	case map[string]interface{}:
		data = append(data, args[0].(map[string]interface{}))

	case []interface{}:
		for cc, item := range args[0].([]interface{}) {
			itemDict, ok := item.(map[string]interface{})
			if !ok {
				msg := fmt.Sprintf("args[0][%d] not a DICT: %s", cc, value.ToStr(item))
				return nil, errors.New(msg)
			}

			if extData != nil {
				for k, v := range extData {
					if _, ok := itemDict[k]; !ok {
						itemDict[k] = v
					}
				}
			}

			data = append(data, itemDict)
		}

	default:
		msg := fmt.Sprintf("args[0] '%s' not a DICT or list of DICT", value.ToStr(args[0]))
		return nil, errors.New(msg)
	}

	smtp, err := emailCheckSmtp(ctx)
	if err != nil {
		return nil, err
	}

	host := smtp["host"].(string)
	port := int(smtp["port"].(int64))
	addr := smtp["address"].(string)
	pswd := smtp["password"].(string)

	msg, err := emailFormatMessage(addr, data, ctx)
	if err != nil {
		return nil, err
	}

	async, _ := ctx.GetXBool("__alert_async__")
	if async {
		go __emailSend(host, port, addr, pswd, msg)
		return nil, err
	}

	return nil, email.NewPlainDialer(host, port, addr, pswd).DialAndSend(msg)
}
