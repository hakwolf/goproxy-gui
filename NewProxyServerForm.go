// 由res2go自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TNewProxyServerForm struct {
    *vcl.TForm
    GroupBox1              *vcl.TGroupBox
    Label1                 *vcl.TLabel
    IpInput                *vcl.TEdit
    Label2                 *vcl.TLabel
    PortInput              *vcl.TEdit
    Label3                 *vcl.TLabel
    ServerProtocolComboBox *vcl.TComboBox
    GroupBox2              *vcl.TGroupBox
    Label12                *vcl.TLabel
    LocalIp                *vcl.TEdit
    Label13                *vcl.TLabel
    LocalPort              *vcl.TEdit
    Label14                *vcl.TLabel
    LocalProtocolComboBox  *vcl.TComboBox
    GroupBox3              *vcl.TGroupBox
    Label10                *vcl.TLabel
    DnsInput               *vcl.TEdit
    Label11                *vcl.TLabel
    TTLInput               *vcl.TEdit
    GroupBox4              *vcl.TGroupBox
    ListView1              *vcl.TListView
    ConfirmBtn             *vcl.TButton
    CancelBtn              *vcl.TButton
    Label15                *vcl.TLabel
    ProxyTypeComboBox      *vcl.TComboBox
    PageControl1           *vcl.TPageControl
    TlsPage                *vcl.TTabSheet
    CrtLabel               *vcl.TLabel
    CrtInput               *vcl.TEdit
    KeyLabel               *vcl.TLabel
    KeyInput               *vcl.TEdit
    SelectCrtBtn           *vcl.TButton
    SelectKeyBtn           *vcl.TButton
    SSHPage                *vcl.TTabSheet
    SshUserNameLabel       *vcl.TLabel
    SshUserNameInput       *vcl.TEdit
    SshPwdLabel            *vcl.TLabel
    SshPwdInput            *vcl.TEdit
    SshKeyLabel            *vcl.TLabel
    SshKeyInput            *vcl.TEdit
    SelectPrivateKeyBtn    *vcl.TButton
    KcpPage                *vcl.TTabSheet
    KcpPwdLabel            *vcl.TLabel
    KcpPwdInput            *vcl.TEdit
    Label4                 *vcl.TLabel
    ServerNameInput        *vcl.TEdit
    PopupMenu1             *vcl.TPopupMenu
    AddMapping             *vcl.TMenuItem
    ModifyMapping          *vcl.TMenuItem
    DelMapping             *vcl.TMenuItem
    ActionList1            *vcl.TActionList
    SelectCrtAction        *vcl.TAction
    SelectKeyAction        *vcl.TAction
    SelectPrivateKeyAction *vcl.TAction
    SelectCrtDialog        *vcl.TOpenDialog
    SelectKeyDialog        *vcl.TOpenDialog
    SelectPrivateKeyDialog *vcl.TOpenDialog

    //::private::
    TNewProxyServerFormFields
}

var NewProxyServerForm *TNewProxyServerForm




// 以字节形式加载
// vcl.Application.CreateForm(newProxyServerFormBytes, &NewProxyServerForm)

var newProxyServerFormBytes = []byte("\x54\x50\x46\x30\x13\x54\x4E\x65\x77\x50\x72\x6F\x78\x79\x53\x65\x72\x76\x65\x72\x46\x6F\x72\x6D\x12\x4E\x65\x77\x50\x72\x6F\x78\x79\x53\x65\x72\x76\x65\x72\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x03\xFC\x01\x06\x48\x65\x69\x67\x68\x74\x03\x31\x02\x03\x54\x6F\x70\x03\xB7\x00\x05\x57\x69\x64\x74\x68\x03\x95\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x96\xB0\xE5\xA2\x9E\xE6\x9C\x8D\xE5\x8A\xA1\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x31\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x95\x02\x0D\x44\x65\x73\x69\x67\x6E\x54\x69\x6D\x65\x50\x50\x49\x02\x78\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x07\x32\x2E\x30\x2E\x32\x2E\x30\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x31\x04\x4C\x65\x66\x74\x02\x10\x06\x48\x65\x69\x67\x68\x74\x03\xB8\x00\x03\x54\x6F\x70\x02\x30\x05\x57\x69\x64\x74\x68\x03\x80\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0F\xE6\x9C\x8D\xE5\x8A\xA1\xE5\x99\xA8\xE9\x85\x8D\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x9F\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x7C\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x10\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x9C\xB0\xE5\x9D\x80\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x07\x49\x70\x49\x6E\x70\x75\x74\x04\x4C\x65\x66\x74\x02\x48\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\xC8\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x03\x28\x01\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xAB\xAF\xE5\x8F\xA3\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x09\x50\x6F\x72\x74\x49\x6E\x70\x75\x74\x04\x4C\x65\x66\x74\x03\x60\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x64\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x33\x04\x4C\x65\x66\x74\x02\x0C\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x38\x05\x57\x69\x64\x74\x68\x02\x22\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x07\x20\xE5\x8D\x8F\xE8\xAE\xAE\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x16\x53\x65\x72\x76\x65\x72\x50\x72\x6F\x74\x6F\x63\x6F\x6C\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x04\x4C\x65\x66\x74\x02\x48\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x30\x05\x57\x69\x64\x74\x68\x02\x7D\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x14\x09\x49\x74\x65\x6D\x49\x6E\x64\x65\x78\x02\x00\x0D\x49\x74\x65\x6D\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x03\x74\x63\x70\x06\x03\x74\x6C\x73\x06\x03\x6B\x63\x70\x06\x03\x73\x73\x68\x06\x02\x77\x73\x06\x03\x77\x73\x73\x00\x08\x4F\x6E\x43\x68\x61\x6E\x67\x65\x07\x1C\x53\x65\x72\x76\x65\x72\x50\x72\x6F\x74\x6F\x63\x6F\x6C\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x43\x68\x61\x6E\x67\x65\x05\x53\x74\x79\x6C\x65\x07\x0E\x63\x73\x44\x72\x6F\x70\x44\x6F\x77\x6E\x4C\x69\x73\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x04\x54\x65\x78\x74\x06\x03\x74\x63\x70\x00\x00\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x32\x04\x4C\x65\x66\x74\x02\x10\x06\x48\x65\x69\x67\x68\x74\x02\x78\x03\x54\x6F\x70\x03\xE8\x00\x05\x57\x69\x64\x74\x68\x03\x80\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x9C\xAC\xE5\x9C\xB0\xE9\x85\x8D\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x5F\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x7C\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x06\x54\x4C\x61\x62\x65\x6C\x07\x4C\x61\x62\x65\x6C\x31\x32\x04\x4C\x65\x66\x74\x02\x0C\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x9C\xB0\xE5\x9D\x80\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x07\x4C\x6F\x63\x61\x6C\x49\x70\x04\x4C\x65\x66\x74\x02\x48\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xC8\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x04\x54\x65\x78\x74\x06\x09\x31\x32\x37\x2E\x30\x2E\x30\x2E\x31\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x07\x4C\x61\x62\x65\x6C\x31\x33\x04\x4C\x65\x66\x74\x03\x40\x01\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xAB\xAF\xE5\x8F\xA3\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x09\x4C\x6F\x63\x61\x6C\x50\x6F\x72\x74\x04\x4C\x65\x66\x74\x03\x70\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x02\x64\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x04\x54\x65\x78\x74\x06\x04\x31\x30\x38\x30\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x07\x4C\x61\x62\x65\x6C\x31\x34\x04\x4C\x65\x66\x74\x02\x0C\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x30\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x8D\x8F\xE8\xAE\xAE\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x15\x4C\x6F\x63\x61\x6C\x50\x72\x6F\x74\x6F\x63\x6F\x6C\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x04\x4C\x65\x66\x74\x02\x48\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x28\x05\x57\x69\x64\x74\x68\x02\x7D\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x14\x09\x49\x74\x65\x6D\x49\x6E\x64\x65\x78\x02\x01\x0D\x49\x74\x65\x6D\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x03\x74\x6C\x73\x06\x03\x74\x63\x70\x06\x03\x6B\x63\x70\x00\x08\x4F\x6E\x43\x68\x61\x6E\x67\x65\x07\x1B\x4C\x6F\x63\x61\x6C\x50\x72\x6F\x74\x6F\x63\x6F\x6C\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x43\x68\x61\x6E\x67\x65\x05\x53\x74\x79\x6C\x65\x07\x0E\x63\x73\x44\x72\x6F\x70\x44\x6F\x77\x6E\x4C\x69\x73\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x04\x54\x65\x78\x74\x06\x03\x74\x63\x70\x00\x00\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x33\x04\x4C\x65\x66\x74\x02\x10\x06\x48\x65\x69\x67\x68\x74\x02\x58\x03\x54\x6F\x70\x03\x60\x01\x05\x57\x69\x64\x74\x68\x03\x80\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE5\x85\xB6\xE4\xBB\x96\xE9\x85\x8D\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x3F\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x7C\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x06\x54\x4C\x61\x62\x65\x6C\x07\x4C\x61\x62\x65\x6C\x31\x30\x04\x4C\x65\x66\x74\x02\x0C\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x11\x05\x57\x69\x64\x74\x68\x02\x20\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\x44\x4E\x53\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x08\x44\x6E\x73\x49\x6E\x70\x75\x74\x04\x4C\x65\x66\x74\x02\x40\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x09\x05\x57\x69\x64\x74\x68\x03\xB8\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x04\x54\x65\x78\x74\x06\x0A\x38\x2E\x38\x2E\x38\x2E\x38\x3A\x35\x33\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x07\x4C\x61\x62\x65\x6C\x31\x31\x04\x4C\x65\x66\x74\x03\x10\x01\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x11\x05\x57\x69\x64\x74\x68\x02\x33\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x08\x54\x54\x4C\x28\xE7\xA7\x92\x29\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x08\x54\x54\x4C\x49\x6E\x70\x75\x74\x04\x4C\x65\x66\x74\x03\x58\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x09\x05\x57\x69\x64\x74\x68\x02\x64\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x04\x54\x65\x78\x74\x06\x03\x36\x30\x30\x00\x00\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x09\x47\x72\x6F\x75\x70\x42\x6F\x78\x34\x04\x4C\x65\x66\x74\x02\x10\x06\x48\x65\x69\x67\x68\x74\x03\x08\x01\x03\x54\x6F\x70\x03\xE8\x00\x05\x57\x69\x64\x74\x68\x03\x80\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE7\xA9\xBF\xE9\x80\x8F\xE6\x98\xA0\xE5\xB0\x84\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xEF\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x7C\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x07\x56\x69\x73\x69\x62\x6C\x65\x08\x00\x09\x54\x4C\x69\x73\x74\x56\x69\x65\x77\x09\x4C\x69\x73\x74\x56\x69\x65\x77\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\xE8\x00\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x78\x02\x07\x43\x6F\x6C\x75\x6D\x6E\x73\x0E\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xB1\xBB\xE5\x9E\x8B\x05\x57\x69\x64\x74\x68\x02\x64\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xBF\x9C\xE7\xA8\x8B\xE7\xAB\xAF\xE5\x8F\xA3\x05\x57\x69\x64\x74\x68\x02\x64\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x9C\xAC\xE5\x9C\xB0\xE7\xAB\xAF\xE5\x8F\xA3\x05\x57\x69\x64\x74\x68\x02\x64\x00\x00\x09\x50\x6F\x70\x75\x70\x4D\x65\x6E\x75\x07\x0A\x50\x6F\x70\x75\x70\x4D\x65\x6E\x75\x31\x09\x52\x6F\x77\x53\x65\x6C\x65\x63\x74\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x09\x56\x69\x65\x77\x53\x74\x79\x6C\x65\x07\x08\x76\x73\x52\x65\x70\x6F\x72\x74\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0A\x43\x6F\x6E\x66\x69\x72\x6D\x42\x74\x6E\x04\x4C\x65\x66\x74\x03\xB0\x00\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x03\x08\x02\x05\x57\x69\x64\x74\x68\x02\x5E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xA1\xAE\xE5\xAE\x9A\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0F\x43\x6F\x6E\x66\x69\x72\x6D\x42\x74\x6E\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x09\x43\x61\x6E\x63\x65\x6C\x42\x74\x6E\x04\x4C\x65\x66\x74\x03\x70\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x03\x08\x02\x05\x57\x69\x64\x74\x68\x02\x5E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x8F\x96\xE6\xB6\x88\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0E\x43\x61\x6E\x63\x65\x6C\x42\x74\x6E\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x07\x4C\x61\x62\x65\x6C\x31\x35\x04\x4C\x65\x66\x74\x02\x10\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x3C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE4\xBB\xA3\xE7\x90\x86\xE7\xB1\xBB\xE5\x9E\x8B\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x11\x50\x72\x6F\x78\x79\x54\x79\x70\x65\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x04\x4C\x65\x66\x74\x02\x58\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x7D\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x14\x09\x49\x74\x65\x6D\x49\x6E\x64\x65\x78\x02\x01\x0D\x49\x74\x65\x6D\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x04\x48\x54\x54\x50\x06\x03\x53\x50\x53\x06\x03\x54\x43\x50\x06\x03\x55\x44\x50\x06\x05\x53\x4F\x43\x4B\x53\x06\x06\xE7\xA9\xBF\xE9\x80\x8F\x00\x08\x4F\x6E\x43\x68\x61\x6E\x67\x65\x07\x17\x50\x72\x6F\x78\x79\x54\x79\x70\x65\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x43\x68\x61\x6E\x67\x65\x05\x53\x74\x79\x6C\x65\x07\x0E\x63\x73\x44\x72\x6F\x70\x44\x6F\x77\x6E\x4C\x69\x73\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x06\x04\x54\x65\x78\x74\x06\x03\x53\x50\x53\x00\x00\x0C\x54\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\x0C\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\x31\x04\x4C\x65\x66\x74\x02\x18\x06\x48\x65\x69\x67\x68\x74\x02\x38\x03\x54\x6F\x70\x03\xA0\x00\x05\x57\x69\x64\x74\x68\x03\x58\x02\x0A\x41\x63\x74\x69\x76\x65\x50\x61\x67\x65\x07\x07\x4B\x63\x70\x50\x61\x67\x65\x08\x53\x68\x6F\x77\x54\x61\x62\x73\x08\x08\x54\x61\x62\x49\x6E\x64\x65\x78\x02\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x07\x07\x56\x69\x73\x69\x62\x6C\x65\x08\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x07\x54\x6C\x73\x50\x61\x67\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x07\x54\x6C\x73\x50\x61\x67\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x30\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x50\x02\x00\x06\x54\x4C\x61\x62\x65\x6C\x08\x43\x72\x74\x4C\x61\x62\x65\x6C\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE8\xAF\x81\xE4\xB9\xA6\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x08\x43\x72\x74\x49\x6E\x70\x75\x74\x04\x4C\x65\x66\x74\x02\x30\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\xC0\x00\x08\x52\x65\x61\x64\x4F\x6E\x6C\x79\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x08\x4B\x65\x79\x4C\x61\x62\x65\x6C\x04\x4C\x65\x66\x74\x03\x30\x01\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\xAF\x86\xE9\x92\xA5\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x08\x4B\x65\x79\x49\x6E\x70\x75\x74\x04\x4C\x65\x66\x74\x03\x50\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\xC4\x00\x08\x52\x65\x61\x64\x4F\x6E\x6C\x79\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0C\x53\x65\x6C\x65\x63\x74\x43\x72\x74\x42\x74\x6E\x04\x4C\x65\x66\x74\x03\xF0\x00\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x18\x06\x41\x63\x74\x69\x6F\x6E\x07\x0F\x53\x65\x6C\x65\x63\x74\x43\x72\x74\x41\x63\x74\x69\x6F\x6E\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x11\x53\x65\x6C\x65\x63\x74\x43\x72\x74\x42\x74\x6E\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0C\x53\x65\x6C\x65\x63\x74\x4B\x65\x79\x42\x74\x6E\x04\x4C\x65\x66\x74\x03\x10\x02\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x18\x06\x41\x63\x74\x69\x6F\x6E\x07\x0F\x53\x65\x6C\x65\x63\x74\x4B\x65\x79\x41\x63\x74\x69\x6F\x6E\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x07\x53\x53\x48\x50\x61\x67\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x07\x53\x53\x48\x50\x61\x67\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x30\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x50\x02\x00\x06\x54\x4C\x61\x62\x65\x6C\x10\x53\x73\x68\x55\x73\x65\x72\x4E\x61\x6D\x65\x4C\x61\x62\x65\x6C\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x2D\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE7\x94\xA8\xE6\x88\xB7\xE5\x90\x8D\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x10\x53\x73\x68\x55\x73\x65\x72\x4E\x61\x6D\x65\x49\x6E\x70\x75\x74\x04\x4C\x65\x66\x74\x02\x47\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x64\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x0B\x53\x73\x68\x50\x77\x64\x4C\x61\x62\x65\x6C\x04\x4C\x65\x66\x74\x03\xDF\x00\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\xAF\x86\xE7\xA0\x81\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0B\x53\x73\x68\x50\x77\x64\x49\x6E\x70\x75\x74\x04\x4C\x65\x66\x74\x03\x13\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x64\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x0B\x53\x73\x68\x4B\x65\x79\x4C\x61\x62\x65\x6C\x04\x4C\x65\x66\x74\x03\xAF\x01\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xA7\x81\xE9\x92\xA5\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0B\x53\x73\x68\x4B\x65\x79\x49\x6E\x70\x75\x74\x04\x4C\x65\x66\x74\x03\xD0\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x64\x08\x52\x65\x61\x64\x4F\x6E\x6C\x79\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x13\x53\x65\x6C\x65\x63\x74\x50\x72\x69\x76\x61\x74\x65\x4B\x65\x79\x42\x74\x6E\x04\x4C\x65\x66\x74\x03\x2F\x02\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x18\x06\x41\x63\x74\x69\x6F\x6E\x07\x16\x53\x65\x6C\x65\x63\x74\x50\x72\x69\x76\x61\x74\x65\x4B\x65\x79\x41\x63\x74\x69\x6F\x6E\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x07\x4B\x63\x70\x50\x61\x67\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x07\x4B\x63\x70\x50\x61\x67\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x30\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x50\x02\x00\x06\x54\x4C\x61\x62\x65\x6C\x0B\x4B\x63\x70\x50\x77\x64\x4C\x61\x62\x65\x6C\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\xAF\x86\xE9\x92\xA5\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0B\x4B\x63\x70\x50\x77\x64\x49\x6E\x70\x75\x74\x04\x4C\x65\x66\x74\x02\x40\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x64\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x34\x04\x4C\x65\x66\x74\x03\x20\x01\x06\x48\x65\x69\x67\x68\x74\x02\x14\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x1E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x90\x8D\xE7\xA7\xB0\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0F\x53\x65\x72\x76\x65\x72\x4E\x61\x6D\x65\x49\x6E\x70\x75\x74\x04\x4C\x65\x66\x74\x03\x68\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1C\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\xC8\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x08\x00\x00\x0A\x54\x50\x6F\x70\x75\x70\x4D\x65\x6E\x75\x0A\x50\x6F\x70\x75\x70\x4D\x65\x6E\x75\x31\x04\x6C\x65\x66\x74\x03\xF0\x01\x03\x74\x6F\x70\x03\x08\x02\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0A\x41\x64\x64\x4D\x61\x70\x70\x69\x6E\x67\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x96\xB0\xE5\xA2\x9E\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0F\x41\x64\x64\x4D\x61\x70\x70\x69\x6E\x67\x43\x6C\x69\x63\x6B\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0D\x4D\x6F\x64\x69\x66\x79\x4D\x61\x70\x70\x69\x6E\x67\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE4\xBF\xAE\xE6\x94\xB9\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x12\x4D\x6F\x64\x69\x66\x79\x4D\x61\x70\x70\x69\x6E\x67\x43\x6C\x69\x63\x6B\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0A\x44\x65\x6C\x4D\x61\x70\x70\x69\x6E\x67\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x88\xA0\xE9\x99\xA4\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0F\x44\x65\x6C\x4D\x61\x70\x70\x69\x6E\x67\x43\x6C\x69\x63\x6B\x00\x00\x00\x0B\x54\x41\x63\x74\x69\x6F\x6E\x4C\x69\x73\x74\x0B\x41\x63\x74\x69\x6F\x6E\x4C\x69\x73\x74\x31\x04\x6C\x65\x66\x74\x03\x70\x02\x03\x74\x6F\x70\x03\x08\x02\x00\x07\x54\x41\x63\x74\x69\x6F\x6E\x0F\x53\x65\x6C\x65\x63\x74\x43\x72\x74\x41\x63\x74\x69\x6F\x6E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\x2E\x2E\x2E\x09\x4F\x6E\x45\x78\x65\x63\x75\x74\x65\x07\x16\x53\x65\x6C\x65\x63\x74\x43\x72\x74\x41\x63\x74\x69\x6F\x6E\x45\x78\x65\x63\x75\x74\x65\x00\x00\x07\x54\x41\x63\x74\x69\x6F\x6E\x0F\x53\x65\x6C\x65\x63\x74\x4B\x65\x79\x41\x63\x74\x69\x6F\x6E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\x2E\x2E\x2E\x09\x4F\x6E\x45\x78\x65\x63\x75\x74\x65\x07\x16\x53\x65\x6C\x65\x63\x74\x4B\x65\x79\x41\x63\x74\x69\x6F\x6E\x45\x78\x65\x63\x75\x74\x65\x00\x00\x07\x54\x41\x63\x74\x69\x6F\x6E\x16\x53\x65\x6C\x65\x63\x74\x50\x72\x69\x76\x61\x74\x65\x4B\x65\x79\x41\x63\x74\x69\x6F\x6E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\x2E\x2E\x2E\x09\x4F\x6E\x45\x78\x65\x63\x75\x74\x65\x07\x1D\x53\x65\x6C\x65\x63\x74\x50\x72\x69\x76\x61\x74\x65\x4B\x65\x79\x41\x63\x74\x69\x6F\x6E\x45\x78\x65\x63\x75\x74\x65\x00\x00\x00\x0B\x54\x4F\x70\x65\x6E\x44\x69\x61\x6C\x6F\x67\x0F\x53\x65\x6C\x65\x63\x74\x43\x72\x74\x44\x69\x61\x6C\x6F\x67\x04\x6C\x65\x66\x74\x03\x10\x02\x03\x74\x6F\x70\x03\x08\x02\x00\x00\x0B\x54\x4F\x70\x65\x6E\x44\x69\x61\x6C\x6F\x67\x0F\x53\x65\x6C\x65\x63\x74\x4B\x65\x79\x44\x69\x61\x6C\x6F\x67\x04\x6C\x65\x66\x74\x03\x30\x02\x03\x74\x6F\x70\x03\x08\x02\x00\x00\x0B\x54\x4F\x70\x65\x6E\x44\x69\x61\x6C\x6F\x67\x16\x53\x65\x6C\x65\x63\x74\x50\x72\x69\x76\x61\x74\x65\x4B\x65\x79\x44\x69\x61\x6C\x6F\x67\x04\x6C\x65\x66\x74\x03\x50\x02\x03\x74\x6F\x70\x03\x08\x02\x00\x00\x00")
