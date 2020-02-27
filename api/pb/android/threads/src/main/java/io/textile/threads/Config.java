package io.textile.threads;

import java.util.function.Consumer;

import io.grpc.ManagedChannel;

public interface Config {
    String getSession();
    void setSession(String session);
    ManagedChannel getChannel();
    void init(Consumer<Boolean> ready);
}
