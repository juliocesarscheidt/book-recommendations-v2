import json
import os
# pip install pyyaml
import yaml
import requests

SVC_NAME = os.environ.get('SVC_NAME')
ENV_NAME = os.environ.get('ENV_NAME')
SVC_PORT = os.environ.get('SVC_PORT', 4080)

API_TOKEN = os.environ.get('API_TOKEN', '')

def make_request(url, method='GET', data=None):
  auth = None
  headers = {
    'Content-Type': 'application/json;charset=UTF-8',
  }
  headers.update({
    'Authorization': 'Bearer ' + API_TOKEN,
  })
  response = requests.request(method=method, json=data, url=url, headers=headers, auth=auth)
  return response.json()

def create_service(svc_name, env_name, svc_port):
  svc_full_name = f'{svc_name}-{env_name}'

  payload = {
    'name': svc_full_name,
    'host': f'{svc_name}.patio-dss-{env_name}.svc.cluster.local',
    'protocol': 'http',
    'port': svc_port,
    'retries': 5,
    'connect_timeout': 60000,
    'write_timeout': 60000,
    'read_timeout': 60000,
    'tags': [svc_name]
  }
  print('payload', payload)

  url = 'http://kong-proxy/services'
  make_request(url, 'POST', payload)

  services = make_request(url, 'GET')
  # with open('services.json') as f:
  #   services = json.loads(f.read())

  svc_id = None

  for service in services['data']:
    # print('service', service)
    if service['name'] == svc_full_name:
      svc_id = service['id']

  print('svc_id', svc_id)

  return svc_id

def create_route(svc_name, env_name, svc_id, route_name, route_path, methods):
  svc_route_full_name = f'{svc_name}.{route_name}'

  payload = {
    'name': svc_route_full_name,
    'hosts': [f'api-{env_name}.oci.ti.rumolog.com'],
    'protocols': ['http', 'https'],
    'methods': methods,
    'paths': [route_path],
    'path_handling': 'v1',
    'strip_path': False,
    'preserve_host': False,
    'https_redirect_status_code': 426,
    'regex_priority': 0,
    'service': {'id': svc_id},
    'tags': [svc_name]
  }
  print('payload', payload)

  url = f'http://kong-proxy/services/{svc_id}/routes'
  make_request(url, 'POST', payload)

  url = 'http://kong-proxy/routes'
  routes = make_request(url, 'GET')
  # with open('routes.json') as f:
  #   routes = json.loads(f.read())

  route_id = None

  for route in routes['data']:
    # print('route', route)
    if route['name'] == svc_route_full_name:
      route_id = route['id']

  print('route_id', route_id)

  # add plugins
  jwt_payload = {
    'name': 'jwt-keycloak',
    'route': {'id': route_id},
    'config': {
      'uri_param_names': ['jwt'],
      'cookie_names': ['oauthtoken'],
      'claims_to_verify': ['exp'],
      'run_on_preflight': True,
      'algorithm': 'RS256',
      'allowed_iss': [f'https://sso.ops-qas.ti.rumolog.com/auth/realms/tlg-{env_name}'],
      'iss_key_grace_period': 10,
      'well_known_template': '%s/.well-known/openid-configuration',
      'consumer_match_claim': 'azp'
    }
  }
  # print('jwt_payload', jwt_payload)

  url = f'http://kong-proxy/plugins'
  make_request(url, 'POST', jwt_payload)

  cors_payload = {
    'enabled': True,
    'protocols': ['grpc', 'grpcs', 'http', 'https'],
    'name': 'cors',
    'route': {'id': route_id},
    'config': {
      'origins': ['*'],
      'headers': ['authorization', 'perfilnome', 'usuariologin', 'usuarionome', 'access-control-allow-origin', 'accept', 'content-type', 'referer', 'user-agent', 'accept-language', 'content-language', 'traceparent'],
      'methods': ['GET', 'POST', 'PUT', 'DELETE', 'OPTIONS', 'PATCH'],
      'credentials': True,
      'preflight_continue': False
    }
  }
  # print('cors_payload', cors_payload)

  url = f'http://kong-proxy/plugins'
  make_request(url, 'POST', cors_payload)

with open(f'etc/openapi/{SVC_NAME}-{ENV_NAME}.yaml') as file:
  swagger = yaml.load(file, Loader=yaml.FullLoader)
  # print(swagger['paths'])

  svc_name = swagger['info']['title']
  # print(svc_name)

  routes_list = []

  routes = list(swagger['paths'].keys())
  # print(routes)

  svc_id = create_service(SVC_NAME, ENV_NAME, SVC_PORT)

  for route in routes:
    # print(swagger['paths'][route])
    for method in list(swagger['paths'][route].keys()):
      route_name = swagger['paths'][route][method]['operationId']
      methods = ['OPTIONS']
      methods.append(method.upper())

      routes_list.append({
        'name': route_name,
        'methods': methods,
        'path': route,
      })

      create_route(SVC_NAME, ENV_NAME, svc_id, route_name, route, methods)

  print(routes_list)
