// Simple ES client base on RESTful interface.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package client

import (
	"bytes"
	"core/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const TIMEOUT int = 60

// ESClient - Structure of ES client.
//
//
type ESClient struct {
	HttpClient *http.Client
}

// Init() - Initialize ESClient instance.
//
// This function return ESClient instance itself for chain operation.
func (this *ESClient) Init() (*ESClient, error) {
	this.HttpClient = &http.Client{Timeout: 5 * time.Second}
	return this, nil
}

// Search() - Search ES server.
//
// @url:  URL string.
// @data: Data to post.
func (this *ESClient) Search(url string, data interface{}) (interface{}, error) {
	bytesData, err := this.formatSearchData(data)
	if err != nil {
		return nil, err
	}

	return this.Post(url+"/_search", bytesData)
}

// IndexBulk() - Index data to ES using _bulk interface.
//
// @url:  URL string.
// @data: Data to post.
func (this *ESClient) IndexBulk(url string, index []byte, data interface{}) (interface{}, error) {
	bytesData, err := this.formatBulkData(index, data)
	if err != nil {
		return nil, err
	}

	return this.Post(url+"/_bulk", bytesData)
}

// Post() - Post data to ES server.
//
// @url:  URL string.
// @data: Data to post.
func (this *ESClient) Post(url string, data []byte) (interface{}, error) {
	rsp, rspError := this.postHttp(url, data)
	if rsp == nil {
		return nil, rspError
	}

	rspValue, err := json.Parse(rsp, "")
	if err != nil {
		msg := fmt.Sprintf("invalid JSON response: %s, %s", err, string(rsp))
		return nil, errors.New(msg)
	}

	if rspError != nil {
		str, _ := json.DumpStr(rspValue)
		msg := fmt.Sprintf("ES server response %s %s", rspError, str)
		return nil, errors.New(msg)
	}

	return rspValue, rspError
}

// formatQueryData() - Format query data to []byte.
//
// @data: Data to post.
func (this *ESClient) formatSearchData(data interface{}) ([]byte, error) {
	var bytesData []byte

	switch data.(type) {
	case []byte:
		bytesData = data.([]byte)

	case string:
		bytesData = []byte(data.(string))

	case map[string]interface{}, []interface{}:
		if ret, err := json.ToBytes(data.(interface{})); err != nil {
			msg := fmt.Sprintf("invalid search data '%s'", err)
			return nil, errors.New(msg)
		} else {
			bytesData = ret
		}

	default:
		return nil, errors.New(fmt.Sprintf("invalid data '%s'", data))
	}

	return bytesData, nil
}

// formatBulkData() - Format query data to []byte.
//
// @data: Data to post.
func (this *ESClient) formatBulkData(index []byte, data interface{}) ([]byte, error) {
	buffer := bytes.Buffer{}

	switch data.(type) {
	case []byte:
		buffer.Write(index)
		buffer.WriteByte('\n')
		buffer.Write(data.([]byte))

	case string:
		buffer.Write(index)
		buffer.WriteByte('\n')
		buffer.WriteString(data.(string))

	case map[string]interface{}:
		buffer.Write(index)
		buffer.WriteByte('\n')
		if bytesData, err := json.ToBytes(data); err != nil {
			msg := fmt.Sprintf("invalid _bulk data '%s'", err)
			return nil, errors.New(msg)
		} else {
			buffer.Write(bytesData)
		}
	case []interface{}:
		for _, item := range data.([]interface{}) {
			buffer.Write(index)
			buffer.WriteByte('\n')

			if bytesData, err := json.ToBytes(item); err != nil {
				msg := fmt.Sprintf("invalid LIST item '%s'", err)
				return nil, errors.New(msg)
			} else {
				buffer.Write(bytesData)
			}
			buffer.WriteByte('\n')
		}

	default:
		return nil, errors.New(fmt.Sprintf("invalid _bulk data '%s'", data))
	}

	return buffer.Bytes(), nil
}

// postHttp() - Post data to URL.
//
// @url:  URL string.
// @data: Byte stream.
func (this *ESClient) postHttp(url string, data []byte) (rsp []byte, err error) {
	var response *http.Response = nil
	var request *http.Request = nil

	for cc := 0; cc < 3; cc++ {
		request, err = http.NewRequest("POST", url, bytes.NewReader(data))
		if err != nil {
			request.Body.Close()
			return nil, err
		}

		if response, err = this.HttpClient.Do(request); err == nil {
			break
		}

		request.Body.Close()
	}

	if err != nil {
		return nil, err
	} else if response == nil {
		return nil, errors.New("POST" + url + ": Got 'nil' response")
	}

	if response.StatusCode != 200 {
		err = errors.New(response.Status)
	}

	rsp, _ = ioutil.ReadAll(response.Body)
	response.Body.Close()
	return rsp, err
}

// SetTimeout() - Set timeout seconds of requests to ES server.
//
// @seconds: Timeout seconds.
func (this *ESClient) SetTimeout(seconds int64) error {
	if seconds < 0 {
		msg := fmt.Sprintf("invalid timeout seconds '%d'", seconds)
		return errors.New(msg)
	}

	this.HttpClient.Timeout = time.Duration(seconds) * time.Second
	return nil
}

// NewESClient() - Create ESClient instance.
func NewESClient() (*ESClient, error) {
	return (&ESClient{}).Init()
}
