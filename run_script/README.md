comand to run .\run_script.exe --f=[file.name] name of runing script file
Example:
.\run_script.exe --f=script.js
RESPONSE:  {"status":"OK","data":{"result":null,"info":"Info: 157676538710174519_123"},"messages":[],"page_debug_id":"","timing":{"before_echo":0.22151398658752441}}

--------------------------------------
SCRIPT SS.INFO:  Info: 157676538710174519_123


to configure the program fill out the following fields in conf.json:
"url": "put here URL for instance run script API",
"headers": [
    {
        "par": "authorization",
        "val": "put here auth token for instance"
    },
    {
        "par": "content-type",
        "val": "application/x-www-form-urlencoded; charset=UTF-8"
    }
]

Settings file example:
{
    "url": "https://bs-dev01.simpleone.ru/v1/ajax-script/run",
    "headers": [
        {
            "par": "authorization",
            "val": "Bearer MM7BF-9pqwIN780EkKdfMm2XKXYwWRtF"
        },
        {
            "par": "content-type",
            "val": "application/x-www-form-urlencoded; charset=UTF-8"
        }
    ]
}