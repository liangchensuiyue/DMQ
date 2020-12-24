package com.liangchen.code;

import com.liangchen.DMQFollowerProto.ClientRegistToFollower;
import com.liangchen.DMQFollowerProto.DMQFollowerServiceGrpc;
import com.liangchen.DMQFollowerProto.MessageData;
import com.liangchen.DMQFollowerProto.Response;
import io.grpc.stub.StreamObserver;

public class protoImpl extends DMQFollowerServiceGrpc.DMQFollowerServiceImplBase {
    @Override
    public StreamObserver<MessageData> clientYieldMsgDataRequest(StreamObserver<Response> responseObserver) {
        return super.clientYieldMsgDataRequest(responseObserver);
    }

    @Override
    public void clientConsumeData(ClientRegistToFollower request, StreamObserver<Response> responseObserver) {
        super.clientConsumeData(request, responseObserver);
    }

    @Override
    public void clientCloseChannel(ClientRegistToFollower request, StreamObserver<Response> responseObserver) {
        super.clientCloseChannel(request, responseObserver);
    }
}
