package com.liangchen.code;

public class DmqBuilder{
    private String host;
    private int port;

    public DmqBuilder(String host, int port) {
        this.host = host;
        this.port = port;
    }

    public  DmqObject buildConsumer(String topic, String group, String key){
        return new DmqObject(this.host,this.port,"consumer",topic,group,key);
    }
    public  DmqObject buildProducer(String topic, String key){
        return new DmqObject(this.host,this.port,"producer",topic,"",key);
    }
}