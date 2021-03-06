package com.liangchen.DMQFollowerProto;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.*;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.*;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.34.0)",
    comments = "Source: follower.proto")
public final class DMQFollowerServiceGrpc {

  private DMQFollowerServiceGrpc() {}

  public static final String SERVICE_NAME = "DMQFollower.DMQFollowerService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<MessageData,
      Response> getClientYieldMsgDataRequestMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ClientYieldMsgDataRequest",
      requestType = MessageData.class,
      responseType = Response.class,
      methodType = io.grpc.MethodDescriptor.MethodType.CLIENT_STREAMING)
  public static io.grpc.MethodDescriptor<MessageData,
      Response> getClientYieldMsgDataRequestMethod() {
    io.grpc.MethodDescriptor<MessageData, Response> getClientYieldMsgDataRequestMethod;
    if ((getClientYieldMsgDataRequestMethod = DMQFollowerServiceGrpc.getClientYieldMsgDataRequestMethod) == null) {
      synchronized (DMQFollowerServiceGrpc.class) {
        if ((getClientYieldMsgDataRequestMethod = DMQFollowerServiceGrpc.getClientYieldMsgDataRequestMethod) == null) {
          DMQFollowerServiceGrpc.getClientYieldMsgDataRequestMethod = getClientYieldMsgDataRequestMethod =
              io.grpc.MethodDescriptor.<MessageData, Response>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.CLIENT_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ClientYieldMsgDataRequest"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  MessageData.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  Response.getDefaultInstance()))
              .setSchemaDescriptor(new DMQFollowerServiceMethodDescriptorSupplier("ClientYieldMsgDataRequest"))
              .build();
        }
      }
    }
    return getClientYieldMsgDataRequestMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.liangchen.DMQFollowerProto.ClientRegistToFollower,
      Response> getClientConsumeDataMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ClientConsumeData",
      requestType = com.liangchen.DMQFollowerProto.ClientRegistToFollower.class,
      responseType = Response.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<com.liangchen.DMQFollowerProto.ClientRegistToFollower,
      Response> getClientConsumeDataMethod() {
    io.grpc.MethodDescriptor<com.liangchen.DMQFollowerProto.ClientRegistToFollower, Response> getClientConsumeDataMethod;
    if ((getClientConsumeDataMethod = DMQFollowerServiceGrpc.getClientConsumeDataMethod) == null) {
      synchronized (DMQFollowerServiceGrpc.class) {
        if ((getClientConsumeDataMethod = DMQFollowerServiceGrpc.getClientConsumeDataMethod) == null) {
          DMQFollowerServiceGrpc.getClientConsumeDataMethod = getClientConsumeDataMethod =
              io.grpc.MethodDescriptor.<com.liangchen.DMQFollowerProto.ClientRegistToFollower, Response>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ClientConsumeData"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.liangchen.DMQFollowerProto.ClientRegistToFollower.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  Response.getDefaultInstance()))
              .setSchemaDescriptor(new DMQFollowerServiceMethodDescriptorSupplier("ClientConsumeData"))
              .build();
        }
      }
    }
    return getClientConsumeDataMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.liangchen.DMQFollowerProto.ClientRegistToFollower,
      Response> getClientCloseChannelMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ClientCloseChannel",
      requestType = com.liangchen.DMQFollowerProto.ClientRegistToFollower.class,
      responseType = Response.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.liangchen.DMQFollowerProto.ClientRegistToFollower,
      Response> getClientCloseChannelMethod() {
    io.grpc.MethodDescriptor<com.liangchen.DMQFollowerProto.ClientRegistToFollower, Response> getClientCloseChannelMethod;
    if ((getClientCloseChannelMethod = DMQFollowerServiceGrpc.getClientCloseChannelMethod) == null) {
      synchronized (DMQFollowerServiceGrpc.class) {
        if ((getClientCloseChannelMethod = DMQFollowerServiceGrpc.getClientCloseChannelMethod) == null) {
          DMQFollowerServiceGrpc.getClientCloseChannelMethod = getClientCloseChannelMethod =
              io.grpc.MethodDescriptor.<com.liangchen.DMQFollowerProto.ClientRegistToFollower, Response>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ClientCloseChannel"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.liangchen.DMQFollowerProto.ClientRegistToFollower.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  Response.getDefaultInstance()))
              .setSchemaDescriptor(new DMQFollowerServiceMethodDescriptorSupplier("ClientCloseChannel"))
              .build();
        }
      }
    }
    return getClientCloseChannelMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static DMQFollowerServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DMQFollowerServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DMQFollowerServiceStub>() {
        @Override
        public DMQFollowerServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DMQFollowerServiceStub(channel, callOptions);
        }
      };
    return DMQFollowerServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static DMQFollowerServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DMQFollowerServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DMQFollowerServiceBlockingStub>() {
        @Override
        public DMQFollowerServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DMQFollowerServiceBlockingStub(channel, callOptions);
        }
      };
    return DMQFollowerServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static DMQFollowerServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DMQFollowerServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DMQFollowerServiceFutureStub>() {
        @Override
        public DMQFollowerServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DMQFollowerServiceFutureStub(channel, callOptions);
        }
      };
    return DMQFollowerServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class DMQFollowerServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 客户端向 follower 发送数据
     * </pre>
     */
    public io.grpc.stub.StreamObserver<MessageData> clientYieldMsgDataRequest(
        io.grpc.stub.StreamObserver<Response> responseObserver) {
      return asyncUnimplementedStreamingCall(getClientYieldMsgDataRequestMethod(), responseObserver);
    }

    /**
     * <pre>
     * 客户端消费数据
     * </pre>
     */
    public void clientConsumeData(com.liangchen.DMQFollowerProto.ClientRegistToFollower request,
        io.grpc.stub.StreamObserver<Response> responseObserver) {
      asyncUnimplementedUnaryCall(getClientConsumeDataMethod(), responseObserver);
    }

    /**
     * <pre>
     * 客户端关闭管道
     * </pre>
     */
    public void clientCloseChannel(com.liangchen.DMQFollowerProto.ClientRegistToFollower request,
        io.grpc.stub.StreamObserver<Response> responseObserver) {
      asyncUnimplementedUnaryCall(getClientCloseChannelMethod(), responseObserver);
    }

    @Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getClientYieldMsgDataRequestMethod(),
            asyncClientStreamingCall(
              new MethodHandlers<
                MessageData,
                Response>(
                  this, METHODID_CLIENT_YIELD_MSG_DATA_REQUEST)))
          .addMethod(
            getClientConsumeDataMethod(),
            asyncServerStreamingCall(
              new MethodHandlers<
                com.liangchen.DMQFollowerProto.ClientRegistToFollower,
                Response>(
                  this, METHODID_CLIENT_CONSUME_DATA)))
          .addMethod(
            getClientCloseChannelMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.liangchen.DMQFollowerProto.ClientRegistToFollower,
                Response>(
                  this, METHODID_CLIENT_CLOSE_CHANNEL)))
          .build();
    }
  }

  /**
   */
  public static final class DMQFollowerServiceStub extends io.grpc.stub.AbstractAsyncStub<DMQFollowerServiceStub> {
    private DMQFollowerServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @Override
    protected DMQFollowerServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DMQFollowerServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * 客户端向 follower 发送数据
     * </pre>
     */
    public io.grpc.stub.StreamObserver<MessageData> clientYieldMsgDataRequest(
        io.grpc.stub.StreamObserver<Response> responseObserver) {
      return asyncClientStreamingCall(
          getChannel().newCall(getClientYieldMsgDataRequestMethod(), getCallOptions()), responseObserver);
    }

    /**
     * <pre>
     * 客户端消费数据
     * </pre>
     */
    public void clientConsumeData(com.liangchen.DMQFollowerProto.ClientRegistToFollower request,
        io.grpc.stub.StreamObserver<Response> responseObserver) {
      asyncServerStreamingCall(
          getChannel().newCall(getClientConsumeDataMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 客户端关闭管道
     * </pre>
     */
    public void clientCloseChannel(com.liangchen.DMQFollowerProto.ClientRegistToFollower request,
        io.grpc.stub.StreamObserver<Response> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getClientCloseChannelMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class DMQFollowerServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<DMQFollowerServiceBlockingStub> {
    private DMQFollowerServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @Override
    protected DMQFollowerServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DMQFollowerServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 客户端消费数据
     * </pre>
     */
    public java.util.Iterator<Response> clientConsumeData(
        com.liangchen.DMQFollowerProto.ClientRegistToFollower request) {
      return blockingServerStreamingCall(
          getChannel(), getClientConsumeDataMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 客户端关闭管道
     * </pre>
     */
    public Response clientCloseChannel(com.liangchen.DMQFollowerProto.ClientRegistToFollower request) {
      return blockingUnaryCall(
          getChannel(), getClientCloseChannelMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class DMQFollowerServiceFutureStub extends io.grpc.stub.AbstractFutureStub<DMQFollowerServiceFutureStub> {
    private DMQFollowerServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @Override
    protected DMQFollowerServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DMQFollowerServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 客户端关闭管道
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<Response> clientCloseChannel(
        com.liangchen.DMQFollowerProto.ClientRegistToFollower request) {
      return futureUnaryCall(
          getChannel().newCall(getClientCloseChannelMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CLIENT_CONSUME_DATA = 0;
  private static final int METHODID_CLIENT_CLOSE_CHANNEL = 1;
  private static final int METHODID_CLIENT_YIELD_MSG_DATA_REQUEST = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final DMQFollowerServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(DMQFollowerServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @Override
    @SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CLIENT_CONSUME_DATA:
          serviceImpl.clientConsumeData((com.liangchen.DMQFollowerProto.ClientRegistToFollower) request,
              (io.grpc.stub.StreamObserver<Response>) responseObserver);
          break;
        case METHODID_CLIENT_CLOSE_CHANNEL:
          serviceImpl.clientCloseChannel((com.liangchen.DMQFollowerProto.ClientRegistToFollower) request,
              (io.grpc.stub.StreamObserver<Response>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @Override
    @SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CLIENT_YIELD_MSG_DATA_REQUEST:
          return (io.grpc.stub.StreamObserver<Req>) serviceImpl.clientYieldMsgDataRequest(
              (io.grpc.stub.StreamObserver<Response>) responseObserver);
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class DMQFollowerServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    DMQFollowerServiceBaseDescriptorSupplier() {}

    @Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return DMQFollowerProto.getDescriptor();
    }

    @Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("DMQFollowerService");
    }
  }

  private static final class DMQFollowerServiceFileDescriptorSupplier
      extends DMQFollowerServiceBaseDescriptorSupplier {
    DMQFollowerServiceFileDescriptorSupplier() {}
  }

  private static final class DMQFollowerServiceMethodDescriptorSupplier
      extends DMQFollowerServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    DMQFollowerServiceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (DMQFollowerServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new DMQFollowerServiceFileDescriptorSupplier())
              .addMethod(getClientYieldMsgDataRequestMethod())
              .addMethod(getClientConsumeDataMethod())
              .addMethod(getClientCloseChannelMethod())
              .build();
        }
      }
    }
    return result;
  }
}
