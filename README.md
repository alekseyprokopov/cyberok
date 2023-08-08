# Тестовое задание для cyberok

__Задача:__

Необходимые функции:
 
(для внешних потребителей)
 
- получить список IP адресов, вернуть в ответ список доменов связанных с этими адресами.
- получить список fqdn, вернуть в ответ список IP адресов
- получить список доменов второго уровня, вернуть для каждого домена whois информацию
 
(для поддержания базы данных в актуальном состоянии)
 
- получить список fqdn и запросить из внешних DNS сервисов IP адреса для этих fqdn и сохранить в бд 
- раз в заданный интервал ( из конфига) обновлять текущую базу адресов из внешних dns серверов
- получить и сохранять whois для доменов второго уровня
 
Дополнительная информация
- Сервис должен корректно отрабатывать сигналы на остановку
- Сервис должен уметь логгировать ключевые моменты своей работы в stdout
- считать что в базе не будет больше 100 000 fqdn

## Build and up

    make up

## Run the app

    make run

# REST API

## Получить список fqdn, вернуть в ответ список IP адресов
### Request

`Post /ip`

    {
        "fqdn_data": ["google.com", "amazon.com"]
    }

### Response

    {
      "status": "OK",
      "fqdn_data": [
        {
            "fqdn": "google.com",
            "ips": [
                "142.250.74.142"
            ]
        },
        {
            "fqdn": "amazon.com",
            "ips": [
                "52.94.236.248",
                "205.251.242.103",
                "54.239.28.85"
            ]
          }
      ]
    }

## Получить список IP адресов, вернуть в ответ список доменов связанных с этими адресами.
### Request

`Post /fqdn`

    {
        "ip_data": ["1.1.1.1", "3.3.3.3"]
    }

### Response

    {
      "status":"OK",
      "fqdn_data": [
        {
          "fqdn":"google.com",
          "ips":["1.1.1.1"]
        },
        {
          "fqdn":"amazon.com",
          "ips":["3.3.3.3"]
        }
      ]
    }

## Получить список доменов второго уровня, вернуть для каждого домена whois информацию.
### Request

`Post /fqdn`

    {
      "domain_data": ["google.com", "amazon.com"]
    }

### Response

    {
        "status": "OK",
        "whois_data": [
            {
                "domain": "amazon.com",
                "info": "   Domain Name: AMAZON.COM\r\n   Registry Domain ID: 281209_DOMAIN_COM-VRSN\r\n   Registrar WHOIS Server: whois.markmonitor.com\r\n   Registrar URL: http://www.markmonitor.com\r\n   Updated Date: 2023-05-16T19:03:14Z\r\n   Creation Date: 1994-11-01T05:00:00Z\r\n   Registry Expiry Date: 2024-10-31T04:00:00Z\r\n   Registrar: MarkMonitor Inc.\r\n   Registrar IANA ID: 292\r\n   Registrar Abuse Contact Email: abusecomplaints@markmonitor.com\r\n   Registrar Abuse Contact Phone: +1.2086851750\r\n   Domain Status: clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited\r\n   Domain Status: clientTransferProhibited https://icann.org/epp#clientTransferProhibited\r\n   Domain Status: clientUpdateProhibited https://icann.org/epp#clientUpdateProhibited\r\n   Domain Status: serverDeleteProhibited https://icann.org/epp#serverDeleteProhibited\r\n   Domain Status: serverTransferProhibited https://icann.org/epp#serverTransferProhibited\r\n   Domain Status: serverUpdateProhibited https://icann.org/epp#serverUpdateProhibited\r\n   Name Server: NS1.AMZNDNS.CO.UK\r\n   Name Server: NS1.AMZNDNS.COM\r\n   Name Server: NS1.AMZNDNS.NET\r\n   Name Server: NS1.AMZNDNS.ORG\r\n   Name Server: NS2.AMZNDNS.CO.UK\r\n   Name Server: NS2.AMZNDNS.COM\r\n   Name Server: NS2.AMZNDNS.NET\r\n   Name Server: NS2.AMZNDNS.ORG\r\n   DNSSEC: unsigned\r\n   URL of the ICANN Whois Inaccuracy Complaint Form: https://www.icann.org/wicf/\r\n>>> Last update of whois database: 2023-08-08T12:04:38Z <<<\r\n\r\nFor more information on Whois status codes, please visit https://icann.org/epp\r\n\r\nNOTICE: The expiration date displayed in this record is the date the\r\nregistrar's sponsorship of the domain name registration in the registry is\r\ncurrently set to expire. This date does not necessarily reflect the expiration\r\ndate of the domain name registrant's agreement with the sponsoring\r\nregistrar.  Users may consult the sponsoring registrar's Whois database to\r\nview the registrar's reported date of expiration for this registration.\r\n\r\nTERMS OF USE: You are not authorized to access or query our Whois\r\ndatabase through the use of electronic processes that are high-volume and\r\nautomated except as reasonably necessary to register domain names or\r\nmodify existing registrations; the Data in VeriSign Global Registry\r\nServices' (\"VeriSign\") Whois database is provided by VeriSign for\r\ninformation purposes only, and to assist persons in obtaining information\r\nabout or related to a domain name registration record. VeriSign does not\r\nguarantee its accuracy. By submitting a Whois query, you agree to abide\r\nby the following terms of use: You agree that you may use this Data only\r\nfor lawful purposes and that under no circumstances will you use this Data\r\nto: (1) allow, enable, or otherwise support the transmission of mass\r\nunsolicited, commercial advertising or solicitations via e-mail, telephone,\r\nor facsimile; or (2) enable high volume, automated, electronic processes\r\nthat apply to VeriSign (or its computer systems). The compilation,\r\nrepackaging, dissemination or other use of this Data is expressly\r\nprohibited without the prior written consent of VeriSign. You agree not to\r\nuse electronic processes that are automated and high-volume to access or\r\nquery the Whois database except as reasonably necessary to register\r\ndomain names or modify existing registrations. VeriSign reserves the right\r\nto restrict your access to the Whois database in its sole discretion to ensure\r\noperational stability.  VeriSign may restrict or terminate your access to the\r\nWhois database for failure to abide by these terms of use. VeriSign\r\nreserves the right to modify these terms at any time.\r\n\r\nThe Registry database contains ONLY .COM, .NET, .EDU domains and\r\nRegistrars.\r\n"
            },
            {
                "domain": "google.com",
                "info": "   Domain Name: GOOGLE.COM\r\n   Registry Domain ID: 2138514_DOMAIN_COM-VRSN\r\n   Registrar WHOIS Server: whois.markmonitor.com\r\n   Registrar URL: http://www.markmonitor.com\r\n   Updated Date: 2019-09-09T15:39:04Z\r\n   Creation Date: 1997-09-15T04:00:00Z\r\n   Registry Expiry Date: 2028-09-14T04:00:00Z\r\n   Registrar: MarkMonitor Inc.\r\n   Registrar IANA ID: 292\r\n   Registrar Abuse Contact Email: abusecomplaints@markmonitor.com\r\n   Registrar Abuse Contact Phone: +1.2086851750\r\n   Domain Status: clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited\r\n   Domain Status: clientTransferProhibited https://icann.org/epp#clientTransferProhibited\r\n   Domain Status: clientUpdateProhibited https://icann.org/epp#clientUpdateProhibited\r\n   Domain Status: serverDeleteProhibited https://icann.org/epp#serverDeleteProhibited\r\n   Domain Status: serverTransferProhibited https://icann.org/epp#serverTransferProhibited\r\n   Domain Status: serverUpdateProhibited https://icann.org/epp#serverUpdateProhibited\r\n   Name Server: NS1.GOOGLE.COM\r\n   Name Server: NS2.GOOGLE.COM\r\n   Name Server: NS3.GOOGLE.COM\r\n   Name Server: NS4.GOOGLE.COM\r\n   DNSSEC: unsigned\r\n   URL of the ICANN Whois Inaccuracy Complaint Form: https://www.icann.org/wicf/\r\n>>> Last update of whois database: 2023-08-08T12:04:38Z <<<\r\n\r\nFor more information on Whois status codes, please visit https://icann.org/epp\r\n\r\nNOTICE: The expiration date displayed in this record is the date the\r\nregistrar's sponsorship of the domain name registration in the registry is\r\ncurrently set to expire. This date does not necessarily reflect the expiration\r\ndate of the domain name registrant's agreement with the sponsoring\r\nregistrar.  Users may consult the sponsoring registrar's Whois database to\r\nview the registrar's reported date of expiration for this registration.\r\n\r\nTERMS OF USE: You are not authorized to access or query our Whois\r\ndatabase through the use of electronic processes that are high-volume and\r\nautomated except as reasonably necessary to register domain names or\r\nmodify existing registrations; the Data in VeriSign Global Registry\r\nServices' (\"VeriSign\") Whois database is provided by VeriSign for\r\ninformation purposes only, and to assist persons in obtaining information\r\nabout or related to a domain name registration record. VeriSign does not\r\nguarantee its accuracy. By submitting a Whois query, you agree to abide\r\nby the following terms of use: You agree that you may use this Data only\r\nfor lawful purposes and that under no circumstances will you use this Data\r\nto: (1) allow, enable, or otherwise support the transmission of mass\r\nunsolicited, commercial advertising or solicitations via e-mail, telephone,\r\nor facsimile; or (2) enable high volume, automated, electronic processes\r\nthat apply to VeriSign (or its computer systems). The compilation,\r\nrepackaging, dissemination or other use of this Data is expressly\r\nprohibited without the prior written consent of VeriSign. You agree not to\r\nuse electronic processes that are automated and high-volume to access or\r\nquery the Whois database except as reasonably necessary to register\r\ndomain names or modify existing registrations. VeriSign reserves the right\r\nto restrict your access to the Whois database in its sole discretion to ensure\r\noperational stability.  VeriSign may restrict or terminate your access to the\r\nWhois database for failure to abide by these terms of use. VeriSign\r\nreserves the right to modify these terms at any time.\r\n\r\nThe Registry database contains ONLY .COM, .NET, .EDU domains and\r\nRegistrars.\r\n"
            }
        ]
    }
    
## Получить список fqdn и запросить из внешних DNS сервисов IP адреса для этих fqdn и сохранить в бд
LookupIp для списка fqdn. Сохранение fqdn и списка ip в бд
### Request

`Post /admin/fqdn`

    {
      "fqdn_data": ["google.com", "amazon.com"]
    }

### Response

    {
      "status":"OK"
    }

## Получить и сохранять whois для доменов второго уровня
запрос whois для списка доментов второго уровня. Сохранение домена второго уровня и whois в бд
### Request

`Post /admin/whois`

    {
      "domain_data": ["google.com", "amazon.com"]
    }

### Response

    {
      "status":"OK"
    }
    
# Поддерживание БД в актуальном состоянии

- получить список fqdn и запросить из внешних DNS сервисов IP адреса для этих fqdn и сохранить в бд (см. REST API, post /admin/fqdn)
- получить и сохранять whois для доменов второго уровня (см. REST API, post /admin/whois)
- раз в заданный интервал ( из конфига) обновлять текущую базу адресов из внешних dns серверов (Реализован updateWorker, который раз в заданный интервал(из конфига.env) обновляет ip адреса для имеющихся в базе fqdn)


# Конфиг (.env) 
    
    #HTTP_SERVER
    HTTP_SERVER_PORT="8000"
    HTTP_SERVER_TIMEOUT=5s
    HTTP_SERVER_IDLE_TIMEOUT=60s
    HTTP_SERVER_STOP_TIMEOUT=10s
    HTTP_EXTERNAL_PORT=8080 #for docker
    
    #DB
    DB_HOST="db"
    DB_PORT="5432"
    DB_NAME="cyberok_db"
    DB_USER="postgres"
    DB_PASSWORD="root"
    DB_SSLMODE="disable"
    DB_EXTERNAL_PORT=5433 #for docker
    
    DB_UPDATE_TIMER="5s"
    
    #MIGRATIONS
    MIGRATION_URL=file://migrations
    
    #DNS
    WHOIS_TIMEOUT=30s
    
    #PG-ADMIN
    PG_ADMIN_PORT="5050"

# БД
### fqdn
#### Миграция

    CREATE TABLE IF NOT EXISTS fqdn
    (
        id   BIGSERIAL NOT NULL PRIMARY KEY,
        name VARCHAR   NOT NULL UNIQUE
    );
    
#### Таблица в pgADmin
![image](https://github.com/alekseyprokopov/cyberok/assets/124125256/ac3e6bf6-0dd5-40d9-ad77-ec3eb797597e)

### ip
#### Миграция

    CREATE TABLE IF NOT EXISTS ip
    (
        id      BIGSERIAL NOT NULL PRIMARY KEY,
        fqdn_id INT REFERENCES fqdn(id),
        ip      VARCHAR NOT NULL
    );
    
#### Таблица в pgADmin
![image](https://github.com/alekseyprokopov/cyberok/assets/124125256/98fffd04-bb61-4476-bfcf-65ebff07a5f9)

### whois
#### Миграция

    CREATE TABLE IF NOT EXISTS whois
    (
        id     BIGSERIAL NOT NULL PRIMARY KEY,
        domain VARCHAR   NOT NULL UNIQUE,
        info   VARCHAR   NOT NULL
    );
    
#### Таблица в pgADmin
![image](https://github.com/alekseyprokopov/cyberok/assets/124125256/3f616450-295a-4172-ae10-3ddec6b1b5ba)

# Дополнительно
- реализованы миграции fqdn, ip, whois
- реализовано логгирование ключевых моментов работы в stdout
- реализован Graceful Shutdown (Interrupt)
- pgAdmin (docker); PGADMIN_DEFAULT_EMAIL: `root@root.com`  PGADMIN_DEFAULT_PASSWORD: `root`
