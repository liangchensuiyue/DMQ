package com.liangchen.code;

import com.liangchen.DMQFollowerProto.Response;

public interface  receiveHandler {
    public abstract void call(Response response);
}
