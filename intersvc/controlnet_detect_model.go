/*
 * @Author: SpenserCai
 * @Date: 2023-08-14 01:00:25
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-20 04:17:09
 * @Description: file content
 */
package intersvc

import (
	SdApiModel "github.com/SpenserCai/sd-webui-go/stablediffusion/models"
)

type ControlnetDetectRequest = SdApiModel.BodyDetectControlnetDetectPost

type ControlnetDetectResponse struct {
	Images []string `json:"images"`
	Info   string   `json:"info"`
}
