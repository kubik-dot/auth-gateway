# utils.py
import os
import logging
import json
import time
from datetime import datetime
from flask import g

logger = logging.getLogger(__name__)

def load_config_from_env():
    config = {}
    for var in ['AUTH_GW_DB_HOST', 'AUTH_GW_DB_PORT', 'AUTH_GW_DB_USER', 'AUTH_GW_DB_PASSWORD', 'AUTH_GW_DB_NAME', 'AUTH_GW_JWT_SECRET']:
        config[var] = os.environ.get(var)
        if not config[var]:
            raise Exception(f"Missing environment variable {var}")
    return config

def format_datetime(dt):
    return dt.strftime('%Y-%m-%d %H:%M:%S')

def json_response(data, status_code=200):
    return json.dumps({'status_code': status_code, 'data': data})

def get_jwt_token():
    return g.get('jwt_token')

def get_user_id():
    return g.get('user_id')

def is_admin():
    user_id = get_user_id()
    return user_id == os.environ.get('AUTH_GW_ADMIN_USER_ID')

def record_access_log(data):
    timestamp = format_datetime(datetime.now())
    log_data = {
        'timestamp': timestamp,
        'request': data.get('request'),
        'response': data.get('response'),
        'user_id': get_user_id()
    }
    logger.info(json.dumps(log_data))