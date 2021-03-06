// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: follower.proto

package com.liangchen.DMQFollowerProto;

/**
 * <pre>
 * 客户端向 follower 节点注册服务
 * </pre>
 *
 * Protobuf type {@code DMQFollower.ClientRegistToFollower}
 */
public final class ClientRegistToFollower extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:DMQFollower.ClientRegistToFollower)
    ClientRegistToFollowerOrBuilder {
private static final long serialVersionUID = 0L;
  // Use ClientRegistToFollower.newBuilder() to construct.
  private ClientRegistToFollower(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private ClientRegistToFollower() {
    nodeid_ = "";
    groupname_ = "";
    topic_ = "";
    key_ = "";
  }

  @Override
  @SuppressWarnings({"unused"})
  protected Object newInstance(
      UnusedPrivateParameter unused) {
    return new ClientRegistToFollower();
  }

  @Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private ClientRegistToFollower(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10: {
            String s = input.readStringRequireUtf8();

            nodeid_ = s;
            break;
          }
          case 18: {
            String s = input.readStringRequireUtf8();

            groupname_ = s;
            break;
          }
          case 26: {
            String s = input.readStringRequireUtf8();

            topic_ = s;
            break;
          }
          case 34: {
            String s = input.readStringRequireUtf8();

            key_ = s;
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return DMQFollowerProto.internal_static_DMQFollower_ClientRegistToFollower_descriptor;
  }

  @Override
  protected FieldAccessorTable
      internalGetFieldAccessorTable() {
    return DMQFollowerProto.internal_static_DMQFollower_ClientRegistToFollower_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            ClientRegistToFollower.class, ClientRegistToFollower.Builder.class);
  }

  public static final int NODEID_FIELD_NUMBER = 1;
  private volatile Object nodeid_;
  /**
   * <pre>
   * 节点id
   * </pre>
   *
   * <code>string nodeid = 1;</code>
   * @return The nodeid.
   */
  @Override
  public String getNodeid() {
    Object ref = nodeid_;
    if (ref instanceof String) {
      return (String) ref;
    } else {
      com.google.protobuf.ByteString bs =
          (com.google.protobuf.ByteString) ref;
      String s = bs.toStringUtf8();
      nodeid_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * 节点id
   * </pre>
   *
   * <code>string nodeid = 1;</code>
   * @return The bytes for nodeid.
   */
  @Override
  public com.google.protobuf.ByteString
      getNodeidBytes() {
    Object ref = nodeid_;
    if (ref instanceof String) {
      com.google.protobuf.ByteString b =
          com.google.protobuf.ByteString.copyFromUtf8(
              (String) ref);
      nodeid_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int GROUPNAME_FIELD_NUMBER = 2;
  private volatile Object groupname_;
  /**
   * <pre>
   * 消费者组名
   * </pre>
   *
   * <code>string groupname = 2;</code>
   * @return The groupname.
   */
  @Override
  public String getGroupname() {
    Object ref = groupname_;
    if (ref instanceof String) {
      return (String) ref;
    } else {
      com.google.protobuf.ByteString bs =
          (com.google.protobuf.ByteString) ref;
      String s = bs.toStringUtf8();
      groupname_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * 消费者组名
   * </pre>
   *
   * <code>string groupname = 2;</code>
   * @return The bytes for groupname.
   */
  @Override
  public com.google.protobuf.ByteString
      getGroupnameBytes() {
    Object ref = groupname_;
    if (ref instanceof String) {
      com.google.protobuf.ByteString b =
          com.google.protobuf.ByteString.copyFromUtf8(
              (String) ref);
      groupname_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int TOPIC_FIELD_NUMBER = 3;
  private volatile Object topic_;
  /**
   * <pre>
   * 话题名称
   * </pre>
   *
   * <code>string topic = 3;</code>
   * @return The topic.
   */
  @Override
  public String getTopic() {
    Object ref = topic_;
    if (ref instanceof String) {
      return (String) ref;
    } else {
      com.google.protobuf.ByteString bs =
          (com.google.protobuf.ByteString) ref;
      String s = bs.toStringUtf8();
      topic_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * 话题名称
   * </pre>
   *
   * <code>string topic = 3;</code>
   * @return The bytes for topic.
   */
  @Override
  public com.google.protobuf.ByteString
      getTopicBytes() {
    Object ref = topic_;
    if (ref instanceof String) {
      com.google.protobuf.ByteString b =
          com.google.protobuf.ByteString.copyFromUtf8(
              (String) ref);
      topic_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int KEY_FIELD_NUMBER = 4;
  private volatile Object key_;
  /**
   * <pre>
   * 密钥
   * </pre>
   *
   * <code>string key = 4;</code>
   * @return The key.
   */
  @Override
  public String getKey() {
    Object ref = key_;
    if (ref instanceof String) {
      return (String) ref;
    } else {
      com.google.protobuf.ByteString bs =
          (com.google.protobuf.ByteString) ref;
      String s = bs.toStringUtf8();
      key_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * 密钥
   * </pre>
   *
   * <code>string key = 4;</code>
   * @return The bytes for key.
   */
  @Override
  public com.google.protobuf.ByteString
      getKeyBytes() {
    Object ref = key_;
    if (ref instanceof String) {
      com.google.protobuf.ByteString b =
          com.google.protobuf.ByteString.copyFromUtf8(
              (String) ref);
      key_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  private byte memoizedIsInitialized = -1;
  @Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!getNodeidBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, nodeid_);
    }
    if (!getGroupnameBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 2, groupname_);
    }
    if (!getTopicBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 3, topic_);
    }
    if (!getKeyBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 4, key_);
    }
    unknownFields.writeTo(output);
  }

  @Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!getNodeidBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, nodeid_);
    }
    if (!getGroupnameBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(2, groupname_);
    }
    if (!getTopicBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, topic_);
    }
    if (!getKeyBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(4, key_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @Override
  public boolean equals(final Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof ClientRegistToFollower)) {
      return super.equals(obj);
    }
    ClientRegistToFollower other = (ClientRegistToFollower) obj;

    if (!getNodeid()
        .equals(other.getNodeid())) return false;
    if (!getGroupname()
        .equals(other.getGroupname())) return false;
    if (!getTopic()
        .equals(other.getTopic())) return false;
    if (!getKey()
        .equals(other.getKey())) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + NODEID_FIELD_NUMBER;
    hash = (53 * hash) + getNodeid().hashCode();
    hash = (37 * hash) + GROUPNAME_FIELD_NUMBER;
    hash = (53 * hash) + getGroupname().hashCode();
    hash = (37 * hash) + TOPIC_FIELD_NUMBER;
    hash = (53 * hash) + getTopic().hashCode();
    hash = (37 * hash) + KEY_FIELD_NUMBER;
    hash = (53 * hash) + getKey().hashCode();
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static ClientRegistToFollower parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static ClientRegistToFollower parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static ClientRegistToFollower parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static ClientRegistToFollower parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static ClientRegistToFollower parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static ClientRegistToFollower parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static ClientRegistToFollower parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static ClientRegistToFollower parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static ClientRegistToFollower parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static ClientRegistToFollower parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static ClientRegistToFollower parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static ClientRegistToFollower parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(ClientRegistToFollower prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @Override
  protected Builder newBuilderForType(
      BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * <pre>
   * 客户端向 follower 节点注册服务
   * </pre>
   *
   * Protobuf type {@code DMQFollower.ClientRegistToFollower}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:DMQFollower.ClientRegistToFollower)
      ClientRegistToFollowerOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return DMQFollowerProto.internal_static_DMQFollower_ClientRegistToFollower_descriptor;
    }

    @Override
    protected FieldAccessorTable
        internalGetFieldAccessorTable() {
      return DMQFollowerProto.internal_static_DMQFollower_ClientRegistToFollower_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              ClientRegistToFollower.class, ClientRegistToFollower.Builder.class);
    }

    // Construct using com.liangchen.DMQFollowerProto.ClientRegistToFollower.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @Override
    public Builder clear() {
      super.clear();
      nodeid_ = "";

      groupname_ = "";

      topic_ = "";

      key_ = "";

      return this;
    }

    @Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return DMQFollowerProto.internal_static_DMQFollower_ClientRegistToFollower_descriptor;
    }

    @Override
    public ClientRegistToFollower getDefaultInstanceForType() {
      return ClientRegistToFollower.getDefaultInstance();
    }

    @Override
    public ClientRegistToFollower build() {
      ClientRegistToFollower result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @Override
    public ClientRegistToFollower buildPartial() {
      ClientRegistToFollower result = new ClientRegistToFollower(this);
      result.nodeid_ = nodeid_;
      result.groupname_ = groupname_;
      result.topic_ = topic_;
      result.key_ = key_;
      onBuilt();
      return result;
    }

    @Override
    public Builder clone() {
      return super.clone();
    }
    @Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        Object value) {
      return super.setField(field, value);
    }
    @Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        Object value) {
      return super.addRepeatedField(field, value);
    }
    @Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof ClientRegistToFollower) {
        return mergeFrom((ClientRegistToFollower)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(ClientRegistToFollower other) {
      if (other == ClientRegistToFollower.getDefaultInstance()) return this;
      if (!other.getNodeid().isEmpty()) {
        nodeid_ = other.nodeid_;
        onChanged();
      }
      if (!other.getGroupname().isEmpty()) {
        groupname_ = other.groupname_;
        onChanged();
      }
      if (!other.getTopic().isEmpty()) {
        topic_ = other.topic_;
        onChanged();
      }
      if (!other.getKey().isEmpty()) {
        key_ = other.key_;
        onChanged();
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @Override
    public final boolean isInitialized() {
      return true;
    }

    @Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      ClientRegistToFollower parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (ClientRegistToFollower) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private Object nodeid_ = "";
    /**
     * <pre>
     * 节点id
     * </pre>
     *
     * <code>string nodeid = 1;</code>
     * @return The nodeid.
     */
    public String getNodeid() {
      Object ref = nodeid_;
      if (!(ref instanceof String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        String s = bs.toStringUtf8();
        nodeid_ = s;
        return s;
      } else {
        return (String) ref;
      }
    }
    /**
     * <pre>
     * 节点id
     * </pre>
     *
     * <code>string nodeid = 1;</code>
     * @return The bytes for nodeid.
     */
    public com.google.protobuf.ByteString
        getNodeidBytes() {
      Object ref = nodeid_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b =
            com.google.protobuf.ByteString.copyFromUtf8(
                (String) ref);
        nodeid_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * 节点id
     * </pre>
     *
     * <code>string nodeid = 1;</code>
     * @param value The nodeid to set.
     * @return This builder for chaining.
     */
    public Builder setNodeid(
        String value) {
      if (value == null) {
    throw new NullPointerException();
  }

      nodeid_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 节点id
     * </pre>
     *
     * <code>string nodeid = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearNodeid() {

      nodeid_ = getDefaultInstance().getNodeid();
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 节点id
     * </pre>
     *
     * <code>string nodeid = 1;</code>
     * @param value The bytes for nodeid to set.
     * @return This builder for chaining.
     */
    public Builder setNodeidBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);

      nodeid_ = value;
      onChanged();
      return this;
    }

    private Object groupname_ = "";
    /**
     * <pre>
     * 消费者组名
     * </pre>
     *
     * <code>string groupname = 2;</code>
     * @return The groupname.
     */
    public String getGroupname() {
      Object ref = groupname_;
      if (!(ref instanceof String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        String s = bs.toStringUtf8();
        groupname_ = s;
        return s;
      } else {
        return (String) ref;
      }
    }
    /**
     * <pre>
     * 消费者组名
     * </pre>
     *
     * <code>string groupname = 2;</code>
     * @return The bytes for groupname.
     */
    public com.google.protobuf.ByteString
        getGroupnameBytes() {
      Object ref = groupname_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b =
            com.google.protobuf.ByteString.copyFromUtf8(
                (String) ref);
        groupname_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * 消费者组名
     * </pre>
     *
     * <code>string groupname = 2;</code>
     * @param value The groupname to set.
     * @return This builder for chaining.
     */
    public Builder setGroupname(
        String value) {
      if (value == null) {
    throw new NullPointerException();
  }

      groupname_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 消费者组名
     * </pre>
     *
     * <code>string groupname = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearGroupname() {

      groupname_ = getDefaultInstance().getGroupname();
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 消费者组名
     * </pre>
     *
     * <code>string groupname = 2;</code>
     * @param value The bytes for groupname to set.
     * @return This builder for chaining.
     */
    public Builder setGroupnameBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);

      groupname_ = value;
      onChanged();
      return this;
    }

    private Object topic_ = "";
    /**
     * <pre>
     * 话题名称
     * </pre>
     *
     * <code>string topic = 3;</code>
     * @return The topic.
     */
    public String getTopic() {
      Object ref = topic_;
      if (!(ref instanceof String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        String s = bs.toStringUtf8();
        topic_ = s;
        return s;
      } else {
        return (String) ref;
      }
    }
    /**
     * <pre>
     * 话题名称
     * </pre>
     *
     * <code>string topic = 3;</code>
     * @return The bytes for topic.
     */
    public com.google.protobuf.ByteString
        getTopicBytes() {
      Object ref = topic_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b =
            com.google.protobuf.ByteString.copyFromUtf8(
                (String) ref);
        topic_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * 话题名称
     * </pre>
     *
     * <code>string topic = 3;</code>
     * @param value The topic to set.
     * @return This builder for chaining.
     */
    public Builder setTopic(
        String value) {
      if (value == null) {
    throw new NullPointerException();
  }

      topic_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 话题名称
     * </pre>
     *
     * <code>string topic = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearTopic() {

      topic_ = getDefaultInstance().getTopic();
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 话题名称
     * </pre>
     *
     * <code>string topic = 3;</code>
     * @param value The bytes for topic to set.
     * @return This builder for chaining.
     */
    public Builder setTopicBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);

      topic_ = value;
      onChanged();
      return this;
    }

    private Object key_ = "";
    /**
     * <pre>
     * 密钥
     * </pre>
     *
     * <code>string key = 4;</code>
     * @return The key.
     */
    public String getKey() {
      Object ref = key_;
      if (!(ref instanceof String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        String s = bs.toStringUtf8();
        key_ = s;
        return s;
      } else {
        return (String) ref;
      }
    }
    /**
     * <pre>
     * 密钥
     * </pre>
     *
     * <code>string key = 4;</code>
     * @return The bytes for key.
     */
    public com.google.protobuf.ByteString
        getKeyBytes() {
      Object ref = key_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b =
            com.google.protobuf.ByteString.copyFromUtf8(
                (String) ref);
        key_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * 密钥
     * </pre>
     *
     * <code>string key = 4;</code>
     * @param value The key to set.
     * @return This builder for chaining.
     */
    public Builder setKey(
        String value) {
      if (value == null) {
    throw new NullPointerException();
  }

      key_ = value;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 密钥
     * </pre>
     *
     * <code>string key = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearKey() {

      key_ = getDefaultInstance().getKey();
      onChanged();
      return this;
    }
    /**
     * <pre>
     * 密钥
     * </pre>
     *
     * <code>string key = 4;</code>
     * @param value The bytes for key to set.
     * @return This builder for chaining.
     */
    public Builder setKeyBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);

      key_ = value;
      onChanged();
      return this;
    }
    @Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:DMQFollower.ClientRegistToFollower)
  }

  // @@protoc_insertion_point(class_scope:DMQFollower.ClientRegistToFollower)
  private static final ClientRegistToFollower DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new ClientRegistToFollower();
  }

  public static ClientRegistToFollower getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<ClientRegistToFollower>
      PARSER = new com.google.protobuf.AbstractParser<ClientRegistToFollower>() {
    @Override
    public ClientRegistToFollower parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new ClientRegistToFollower(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<ClientRegistToFollower> parser() {
    return PARSER;
  }

  @Override
  public com.google.protobuf.Parser<ClientRegistToFollower> getParserForType() {
    return PARSER;
  }

  @Override
  public ClientRegistToFollower getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

