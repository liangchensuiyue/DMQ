package com.liangchen.code;

public class DmqBuilder{
    private String host;
    private int port;

    public DmqBuilder(String host, int port) {
        this.host = host;
        this.port = port;
    }

    public  DmqObject buildConsumer(String topic, String group){
        return new DmqObject(this.host,this.port,"consumer",topic,group);
    }
    public  DmqObject buildProducer(String topic){
        return new DmqObject(this.host,this.port,"producer",topic,"");
    }
}