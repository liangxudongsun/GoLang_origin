package logicservice

import (

	"github.com/duanhf2012/origin/Test/msgpb"
	"github.com/duanhf2012/origin/network"
	"github.com/duanhf2012/origin/sysservice"
	"github.com/golang/protobuf/proto"
	"time"

	"github.com/duanhf2012/origin/originnode"
	"github.com/duanhf2012/origin/service"

)

type SubNet1_Service struct {
	service.BaseService
}

func init() {
	originnode.InitService(&SubNet1_Service{})
}

//OnInit ...
func (ws *SubNet1_Service) OnInit() error {
	sysservice.DefaultTSPbService().RegConnectEvent(ws.ConnEventHandler)
	sysservice.DefaultTSPbService().RegDisconnectEvent(ws.DisconnEventHandler)
	sysservice.DefaultTSPbService().RegExceptMessage(ws.ExceptMessage)
	sysservice.DefaultTSPbService().RegMessage(110,&msgpb.Test{},ws.MessageHandler)

	return nil
}


//OnRun ...
func (ws *SubNet1_Service) OnRun() bool {


	time.Sleep(time.Second*10)
	var cli network.TcpSocketClient
	cli.Connect("127.0.0.1:9004")
	test := msgpb.Test{}
	test.AssistCount = proto.Int32(343)



	cli.SendMsg(110,&test)
	cli.SendMsg(110,&test)
	return false
}


 func (ws *SubNet1_Service) MessageHandler(pClient *network.SClient,msgtype uint16,msg proto.Message){
 }

func (ws *SubNet1_Service) ConnEventHandler (pClient *network.SClient){
}

func (ws *SubNet1_Service) DisconnEventHandler (pClient *network.SClient){
}

func (ws *SubNet1_Service) ExceptMessage(pClient *network.SClient,pPack *network.MsgBasePack,err error){
}
