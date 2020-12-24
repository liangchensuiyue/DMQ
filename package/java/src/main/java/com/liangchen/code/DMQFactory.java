package com.liangchen.code;


public class DMQFactory {
    public static void send(){

    }
    public static DmqBuilder newBuilder(String host,int port){
        return new DmqBuilder(host, port);
    }
}
