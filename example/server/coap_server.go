package main

import (
	"log"
	"net"

	"github.com/dustin/go-coap"
)

func handleAuth(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
	log.Printf("Got message in handleAuth: path=%q: %#v from %v", m.Path(), string(m.Payload), a)
	if m.IsConfirmable() {
		res := &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Created,
			MessageID: m.MessageID,
			Token:     m.Token,
			Payload:   []byte("Token for Coap"),
		}
		res.SetOption(coap.ContentFormat, coap.TextPlain)

		log.Printf("Transmitting from Auth %#v", string(res.Payload))
		return res
	}
	return nil
}

func handleMqttPub(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
	log.Printf("Got message in handleMqttPub: path=%q: %#v from %v", m.Path(), string(m.Payload), a)
	if m.IsConfirmable() {
		res := &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Created,
			MessageID: m.MessageID,
			Token:     m.Token,
			Payload:   []byte("Publish message success"),
		}
		res.SetOption(coap.ContentFormat, coap.TextPlain)

		log.Printf("Transmitting from MqttPub %#v", string(res.Payload))
		return res
	}
	return nil
}

func handleMqttRpc(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
	log.Printf("Got message in handleMqttRpc: path=%q: %#v from %v", m.Path(), string(m.Payload), a)
	if m.IsConfirmable() {
		res := &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Created,
			MessageID: m.MessageID,
			Token:     m.Token,
			Payload:   []byte("Rpc Call success"),
		}
		res.SetOption(coap.ContentFormat, coap.TextPlain)

		log.Printf("Transmitting from MqttRpc %#v", string(res.Payload))
		return res
	}
	return nil
}

func main() {
	mux := coap.NewServeMux()
	mux.Handle("/auth", coap.FuncHandler(handleAuth))
	mux.Handle("/mqttpub", coap.FuncHandler(handleMqttPub))
	mux.Handle("/mqttrpc", coap.FuncHandler(handleMqttRpc))

	log.Fatal(coap.ListenAndServe("udp", "127.0.0.1:5683", mux))
}
