import sys
import mysql.connector

def ConnectDB(addr, port, userName, password, dbname):
    '''
    连接数据库，并返回该连接
    host:数据库地址及端口号，如127.0.0.1:8080
    user:用户名
    pwd:用户密码
    dbname:数据库名称
    '''
    config = {'host':addr, 'user':userName, 'password':password, 
    'port':port, 'database':dbname}
    try:
        conn = mysql.connector.connect(**config)
    except Exception as e:
        if e.errno == mysql.connector.errorcode.ER_ACCESS_DENIED_ERROR:
            print("Something is wrong with your username or password")
        elif e.errno == mysql.connector.errorcode.ER_BAD_DB_ERROR:
            print("Database does not exist")
        else:
            print(e)
        sys.exit()
    return conn

def RunSelectCommand(conn, sql):
    '''
    执行SQL语句
    conn:数据库连接
    sql:要执行的sql语句
    '''
    cur = conn.cursor()
    try:
        cur.execute(sql)
    except Exception as e:
        print(e)
        cur.close()
        sys.exit()
    return cur

def RunInsertCommand(conn, sql):
    '''
    执行SQL语句
    conn:数据库连接
    sql:要执行的sql语句
    '''
    cur = conn.cursor()
    try:
        cur.execute(sql)
        #一定要有这句，否则不能插入
        conn.commit()
        #cur.close()
    except Exception as e:
        print(e)
        cur.close()
        sys.exit()
    return cur

def GetData(cur):
    alldata = cur.fetchall()
    return alldata

def CloseConnect(conn):
    conn.close()


    
    
