import json
import requests

# headers = {"authorization":"Bearer eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJndWVzdCIsInVzZXJSb2xlcyI6WyJST0xFX1VTRVIiXSwiaWF0IjoxNjYzODAyNjkxLCJleHAiOjE2NjM4Mzg2OTF9.gZNe4ZDaPUlHwTJFIPZP_et1k0BPqOFPNqiD8FCGRXtmoAQ9xkCNFtwINpgM5_bwDLuEDPDiB5oOhLG-Sp9aqA"}
url = "https://tw-yk.website:81/works/getAllWorks"
login = "https://tw-yk.website:81/auth/login?userName=guest&userPassword=guest"
r = requests.post(login)
print(r.json())