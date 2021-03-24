// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: follower.proto

package com.liangchen.DMQFollowerProto;

public interface MessageDataOrBuilder extends
    // @@protoc_insertion_point(interface_extends:DMQFollower.MessageData)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * 数据对应的topic
   * </pre>
   *
   * <code>string topic = 1;</code>
   * @return The topic.
   */
  java.lang.String getTopic();
  /**
   * <pre>
   * 数据对应的topic
   * </pre>
   *
   * <code>string topic = 1;</code>
   * @return The bytes for topic.
   */
  com.google.protobuf.ByteString
      getTopicBytes();

  /**
   * <pre>
   * 真正的数据
   * </pre>
   *
   * <code>string message = 2;</code>
   * @return The message.
   */
  java.lang.String getMessage();
  /**
   * <pre>
   * 真正的数据
   * </pre>
   *
   * <code>string message = 2;</code>
   * @return The bytes for message.
   */
  com.google.protobuf.ByteString
      getMessageBytes();

  /**
   * <code>int64 length = 3;</code>
   * @return The length.
   */
  long getLength();

  /**
   * <pre>
   * 附加信息
   * </pre>
   *
   * <code>string des = 4;</code>
   * @return The des.
   */
  java.lang.String getDes();
  /**
   * <pre>
   * 附加信息
   * </pre>
   *
   * <code>string des = 4;</code>
   * @return The bytes for des.
   */
  com.google.protobuf.ByteString
      getDesBytes();

  /**
   * <pre>
   *密钥
   * </pre>
   *
   * <code>string key = 5;</code>
   * @return The key.
   */
  java.lang.String getKey();
  /**
   * <pre>
   *密钥
   * </pre>
   *
   * <code>string key = 5;</code>
   * @return The bytes for key.
   */
  com.google.protobuf.ByteString
      getKeyBytes();
}
