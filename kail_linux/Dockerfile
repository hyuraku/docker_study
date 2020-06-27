FROM kalilinux/kali-rolling:latest
RUN apt update && apt install curl nmap traceroute iputils-ping net-tools postgresql procps -y
RUN curl https://raw.githubusercontent.com/rapid7/metasploit-omnibus/master/config/templates/metasploit-framework-wrappers/msfupdate.erb > msfinstall && chmod 755 msfinstall && ./msfinstall
