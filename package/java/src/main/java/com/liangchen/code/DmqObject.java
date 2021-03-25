package com.liangchen.code;

import io.grpc.ManagedChannel;
import com.liangchen.DMQFollowerProto.*;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;

import java.util.UUID;




public class DmqObject {
    private String type;
    private String host;
    private int    port;
    private String consumerGroup;
    private String topic;
    private String key;
    public receiveHandler receiveCall = null;
    public completeHandler completeCall = null;
    public errorHandler errorCall = null;

    public StreamObserver<MessageData> produce = null;

    public DmqObject(String host, int port,String type,String topic, String consumerGroup, String key) {
        this.type = type;
        this.host = host;
        this.port = port;
        this.consumerGroup = consumerGroup;
        this.topic = topic;
        this.key = key;
    }
    public void receive(receiveHandler handler){
        this.receiveCall = handler;
    }
    public void completed(completeHandler handler){
        this.completeCall = handler;
    }
    public void error(errorHandler handler){
        this.errorCall = handler;
    }
    public void connect() throws Exception {
        ManagedChannel managedChannel = ManagedChannelBuilder.forAddress(host, port).
                usePlaintext().build();

        DMQFollowerServiceGrpc.DMQFollowerServiceStub stub = DMQFollowerServiceGrpc.newStub(managedChannel);



        if (this.type.equals("consumer")){

            StreamObserver<Response> DmqConsumerStreamObserver = new StreamObserver<Response>() {

                @Override
                public void onNext(Response response) {
                    if (DmqObject.this.receiveCall == null){
                        return;
                    }
                    DmqObject.this.receiveCall.call(response);
                }

                @Override
                public void onError(Throwable throwable) {
                    System.out.println(throwable.getMessage());
                    if (DmqObject.this.errorCall == null){
                        return;
                    }
                    DmqObject.this.errorCall.call();
                }

                @Override
                public void onCompleted() {
                    if (DmqObject.this.completeCall == null){
                        return;
                    }
                    DmqObject.this.completeCall.call();
                }
            };
            // 只要客户端以流式数据请求时,那么只能非阻塞(异步)来接收数据(不能用blockingStub了)
            ClientRegistToFollower request = ClientRegistToFollower.newBuilder().setGroupname(this.getConsumerGroup()).setTopic(this.getTopic()).setNodeid(UUID.randomUUID().toString()).setKey(this.getKey()).build();
            stub.clientConsumeData(request,DmqConsumerStreamObserver);

        }else if(this.type.equals("producer")){

            StreamObserver<Response> DmqProducerStreamObserver = new StreamObserver<Response>() {


                @Override
                public void onNext(Response response) {
                    if (DmqObject.this.receiveCall == null){
                        return;
                    }
                    DmqObject.this.receiveCall.call(response);
                }

                @Override
                public void onError(Throwable throwable) {
                    System.out.println(throwable.getMessage());
                    if (DmqObject.this.errorCall == null){
                        return;
                    }
                    DmqObject.this.errorCall.call();
                }

                @Override
                public void onCompleted() {
                    if (DmqObject.this.completeCall == null){
                        return;
                    }
                    DmqObject.this.completeCall.call();
                }
            };
            // 只要客户端以流式数据请求时,那么只能非阻塞(异步)来接收数据(不能用blockingStub了)
            produce = stub.clientYieldMsgDataRequest(DmqProducerStreamObserver);


        }else{
            throw new Exception("不存在的类型: " + this.type);
        }
    }
    public void end() {
        if (this.type.equals("consumer")){
            return;
        }
        if (this.type.equals("producer")){
            this.produce.onCompleted();
        }
    }
    public void send(String message) throws Exception {
//        Thread.sleep(50);
        if (this.type.equals("consumer")){
            throw  new Exception("当前类型是消费者,不能生产数据");
        }
        if (this.type.equals("producer")){
            this.produce.onNext(MessageData.newBuilder().setMessage(message).setTopic(this.topic).setKey(this.key).build());
        }
    }
    public String getType() {
        return type;
    }

    public String getHost() {
        return host;
    }
    public String getKey() {
        return key;
    }

    public int getPort() {
        return port;
    }

    public String getConsumerGroup() {
        return consumerGroup;
    }

    public String getTopic() {
        return topic;
    }

}
