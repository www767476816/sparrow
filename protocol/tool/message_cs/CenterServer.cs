//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

// Generated from: CenterServer.proto
namespace CenterServer
{
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"ReqFreeServerResult")]
  public partial class ReqFreeServerResult : global::ProtoBuf.IExtensible
  {
    public ReqFreeServerResult() {}
    
    private bool _success = default(bool);
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"success", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue(default(bool))]
    public bool success
    {
      get { return _success; }
      set { _success = value; }
    }
    private string _IPAddress = "";
    [global::ProtoBuf.ProtoMember(2, IsRequired = false, Name=@"IPAddress", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string IPAddress
    {
      get { return _IPAddress; }
      set { _IPAddress = value; }
    }
    private int _Port = default(int);
    [global::ProtoBuf.ProtoMember(3, IsRequired = false, Name=@"Port", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(default(int))]
    public int Port
    {
      get { return _Port; }
      set { _Port = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
    [global::ProtoBuf.ProtoContract(Name=@"CenterServerMsg")]
    public enum CenterServerMsg
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"CENTERSERVER_UNKNOWN", Value=0)]
      CENTERSERVER_UNKNOWN = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"CTS_M_SERVICE", Value=1001)]
      CTS_M_SERVICE = 1001,
            
      [global::ProtoBuf.ProtoEnum(Name=@"CTS_S_REQUEST_FREE_LOGIN_SERVER", Value=0)]
      CTS_S_REQUEST_FREE_LOGIN_SERVER = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"CTS_S_REQUEST_FREE_GAME_SERVER", Value=1)]
      CTS_S_REQUEST_FREE_GAME_SERVER = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"STC_M_SERVICE_RESULT", Value=11001)]
      STC_M_SERVICE_RESULT = 11001,
            
      [global::ProtoBuf.ProtoEnum(Name=@"STC_S_REQUEST_FREE_LOGIN_SERVER_RESULT", Value=0)]
      STC_S_REQUEST_FREE_LOGIN_SERVER_RESULT = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"STC_S_REQUEST_FREE_GAME_SERVER_RESULT", Value=1)]
      STC_S_REQUEST_FREE_GAME_SERVER_RESULT = 1
    }
  
}