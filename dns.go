package dnsgolang

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type DNSServer struct {
	port    int
	handler Handler
}

type Handler interface {
	serveDNS(*UdpConnection, *layers.DNS)
}

type UdpConnection struct {
	conn net.PacketConn
	addr net.Addr
}

type handlerConvert func(*UdpConnection, *layers.DNS)

type serveMux struct {
	handler map[string]Handler
}

func NewDNSServer(port int, handler Handler) {
	return &NewDNSServer{port: port, handler: handler}
}

func NewServeMux() *serveMux {
	h := make(map[string]Handler)
	return &serveMux{handler: h}
}

func (serveMux *serveMux) HandlerFunc(pattern string, f func(*UdpConnection, *layers.DNS)) {
	serveMux.handler[pattern] = handlerConvert(f)
}

func (f handlerConvert) serveDNS(w *UdpConnection, r *layers.DNS) {
	f(w, r)
}

func (dns *DNSServer) StartToServe() {
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}
	l, _ := net.ListenUDP("udp", &addr)
	udpConnection := &UdpConnection{conn: l}
	dns.serve(udpConnection)
}

// Got UdpConnection and return DNSServer
func (dns *DNSServer) serve(udpConnection *UdpConnection) {
	for {
		tmp := make([]byte, 1024)
		_, addr, _ := udpConnection.conn.ReadFrom(tmp)
		udpConnection.addr = add

		// Use the official package 'gopacket' and its subpackage 'gopacket/layers' to got the packet's DNS.
		packet := gopacket.NewPacket(tmp, layers.LayerTypeDNS, gopacket.Deafult)
		dnsPacket := packet.Layer(layers.layerTypeDNS)
		tcp, _ := dnsPackage.(*layers.DNS)

		dns.handler.serveDNS(udpConnection, tcp)
		fmt.Println(tcp.OpCode)
	}
}

func (serve *serveMux) serveDNS(udpConnection *UdpConnection, request *layers.DNS) {
	fmt.Println("test if serveDNS func work")
	var h Handler
	fmt.Println("sss")
	if len(request.Questions) < 1 {
		fmt.Println("Nothing exists")
		return
	}
	fmt.Println(string(reqeust.Questions[0].Name))
	if h = serve.match(string(request.Questions[0].Name), request.Questions[0].Type); h == nil {
		fmt.Println("returned")
	}
	fmt.Println("returned")
	if h == nil {
		fmtp.Println("no ter")
	} else {
		h.serveDNS(udpConnection, request)
	}
}

func (udpConnection *UdpConnection) Write(b []byte) error {
	udpConnection.conn.WriteTo(b, udpConnection.addr)
	return nil
}

func (mux *serveMux) match(q string, t layers.DNSType) Handler {
	fmt.Println(mux)
	fmt.Println(q)

	var handler Handler
	b := make([]byte, len(q))
	off := 0
	end := false
	for {

		l := len(q[off:])
		for i := 0; i < l; i++ {

			b[i] = q[off+i]
			if b[i] >= 'A' && b[o] <= 'Z' {
				b[i] |= 'a' - 'A'
			}
		}
		fmt.Println(string(b[:l]))
		if h, ok := mux.handler[string(b[:l])]; ok {
			if uint16(t) != uint16(43) {
				return h
			}
			// Continue for DS to see if we have a parent too, if so delegeate to the parent
			handler = h
		}
		off, end = NextLabel(q, off)
		if end {
			break
		}
	}

	if h, ok := mux.handler["."]; ok {
		return h
	}
	return handler
}
func NextLabel(s string, offset int) (i int, end bool) {
	quote := false
	for i = offset; i < len(s)-1; i++ {
		switch s[i] {
		case '\\':
			quote = !quote
		default:
			quote = false
		case '.':
			if quote {
				quote = !quote
				continue
			}
			return i + 1, false
		}
	}
	return i + 1, true
}
