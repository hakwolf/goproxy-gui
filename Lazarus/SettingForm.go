// 由res2go自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TSettingForm struct {
    *vcl.TForm
    AutoStartChk *vcl.TCheckBox
    SysProxyChk  *vcl.TCheckBox
    PortEdit     *vcl.TEdit
    PortLabel    *vcl.TLabel
    ConfirmBtn   *vcl.TButton
    CancelBtn    *vcl.TButton

    //::private::
    TSettingFormFields
}

var SettingForm *TSettingForm




// 以字节形式加载
// vcl.Application.CreateForm(settingFormBytes, &SettingForm)

var settingFormBytes = []byte("\x54\x50\x46\x30\x0C\x54\x53\x65\x74\x74\x69\x6E\x67\x46\x6F\x72\x6D\x0B\x53\x65\x74\x74\x69\x6E\x67\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x03\x48\x02\x06\x48\x65\x69\x67\x68\x74\x03\xA4\x00\x03\x54\x6F\x70\x03\xDC\x01\x05\x57\x69\x64\x74\x68\x03\xE5\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE7\xB3\xBB\xE7\xBB\x9F\xE8\xAE\xBE\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xA4\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xE5\x01\x0D\x44\x65\x73\x69\x67\x6E\x54\x69\x6D\x65\x50\x50\x49\x02\x78\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x07\x31\x2E\x38\x2E\x34\x2E\x30\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x0C\x41\x75\x74\x6F\x53\x74\x61\x72\x74\x43\x68\x6B\x04\x4C\x65\x66\x74\x02\x37\x06\x48\x65\x69\x67\x68\x74\x02\x18\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x54\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\x87\xAA\xE5\x8A\xA8\xE5\x90\xAF\xE5\x8A\xA8\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x0B\x53\x79\x73\x50\x72\x6F\x78\x79\x43\x68\x6B\x04\x4C\x65\x66\x74\x02\x37\x06\x48\x65\x69\x67\x68\x74\x02\x18\x03\x54\x6F\x70\x02\x40\x05\x57\x69\x64\x74\x68\x02\x54\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE7\xB3\xBB\xE7\xBB\x9F\xE4\xBB\xA3\xE7\x90\x86\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x05\x54\x45\x64\x69\x74\x08\x50\x6F\x72\x74\x45\x64\x69\x74\x04\x4C\x65\x66\x74\x03\xD0\x00\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x3C\x05\x57\x69\x64\x74\x68\x03\x98\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x04\x54\x65\x78\x74\x06\x0E\x31\x32\x37\x2E\x30\x2E\x30\x2E\x31\x3A\x38\x30\x38\x30\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x09\x50\x6F\x72\x74\x4C\x61\x62\x65\x6C\x04\x4C\x65\x66\x74\x03\xA8\x00\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x42\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xAB\xAF\xE5\x8F\xA3\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0A\x43\x6F\x6E\x66\x69\x72\x6D\x42\x74\x6E\x04\x4C\x65\x66\x74\x02\x48\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x02\x73\x05\x57\x69\x64\x74\x68\x02\x5E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xA1\xAE\xE5\xAE\x9A\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0F\x43\x6F\x6E\x66\x69\x72\x6D\x42\x74\x6E\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x09\x43\x61\x6E\x63\x65\x6C\x42\x74\x6E\x04\x4C\x65\x66\x74\x03\x18\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x02\x73\x05\x57\x69\x64\x74\x68\x02\x5E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x8F\x96\xE6\xB6\x88\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0E\x43\x61\x6E\x63\x65\x6C\x42\x74\x6E\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x00")
