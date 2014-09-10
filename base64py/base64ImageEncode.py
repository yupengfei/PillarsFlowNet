import base64

f=open(r"/home/developer/QML/PillarsFlow/qml/PillarsFlow/Pictures/PeopleTest/1.jpg",'rb') #二进制方式打开图文件
ls_f=base64.b64encode(f.read()) #读取文件内容，转换为base64编码 
f.close()
f = open("base64.test", 'w')
f.write(str(ls_f))
f.close()