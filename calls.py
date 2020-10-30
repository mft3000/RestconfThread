import requests

def main():

	for id in range(1, 6):
		r = requests.post('http://localhost:8080/restconf/operations/tailf-ncs:devices/tailf-ncs:device=R' + str(id) + '/tailf-ncs:sync-from', auth=('admin', 'admin'), headers={'Content-Type': 'application/yang-data+json'})
		print( r.text )

main()
