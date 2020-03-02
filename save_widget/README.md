This program is simplify the creation of widgets on the SimpleOne platform.

First, create a widget in the system, save it and display it on the necessary page
after these steps, configure the program in the config.json file
To do this, fill out the following fields:

"widgetRestUrl": "put here REST URL for the created widget",
"widgetInstanceUrl": "put here the URL of the created widget instance",
"auth": {
    "login": "put your login here",
    "pass": "put your password here"
},
"files": {
    "htmlTemplate": "put here HTML template file name",
    "css": "put here CSS file name",
    "serverScript": "put here serverScript file name",
    "clientScript": "put here clientScript file name"
}
create a widget in the specified files and run the programm in the console
.\save_widget.exe
if successful, you will get an answer:
Widget saving:  OK
Opening page:  {widgetInstanceUrl}

Settings file example:
{
    "widgetRestUrl": "https://bs-dev01.simpleone.ru/rest/v1/table/sys_widget/158313163911942189",
    "widgetInstanceUrl": "https://bs-dev01.simpleone.ru/record/testing_",
    "auth": {
        "login": "test.test",
        "pass": "123456"
    },
    "files": {
        "htmlTemplate": "test.html",
        "css": "test.css",
        "serverScript": "serverScript.js",
        "clientScript": "clientScript.js"
    }
}