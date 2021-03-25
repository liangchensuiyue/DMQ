package com.liangchen.test;

//import com.liangchen.code.DMQFactory;
import com.liangchen.DMQFollowerProto.Response;
import com.liangchen.code.DMQFactory;
import com.liangchen.code.DmqBuilder;
import com.liangchen.code.DmqObject;
import org.junit.Test;

import java.io.BufferedInputStream;
import java.io.BufferedReader;
import java.io.FileInputStream;
import java.io.InputStreamReader;
import java.util.Scanner;

public class test {

    public int i = 0;
    @Test
    public void consume() throws Exception {
        final  int i = 0;
        DmqBuilder builder =  DMQFactory.newBuilder("127.0.0.1",9208);

        // 创建一个消费者
        DmqObject  consumer =  builder.buildConsumer("phone", "test","hxBA7bZlcJKE72kZ1adLmdnjhGx0Q+PgM3joSeLxcTM=");
        consumer.receive( response -> {
            test.this.i++;
//            if (test.this.i % 1000 == 0){
                System.out.printf("%d  %s\n",test.this.i,response.getData());
//            }
        });
        // 打开链接
        consumer.connect();
        Thread.sleep(10000000);
    }
    @Test
    public void consume1() throws Exception {
        final  int i = 0;
        DmqBuilder builder =  DMQFactory.newBuilder("127.0.0.1",9202);

        // 创建一个消费者
        DmqObject  consumer =  builder.buildConsumer("default", "test", "hxBA7bZlcJKE72kZ1adLmdnjhGx0Q+PgM3joSeLxcTM=");

        consumer.receive( response -> {
            test.this.i++;
            if (test.this.i % 1000 == 0){
                System.out.printf("%d  %s\n",test.this.i,response.getData());
            }
        });
        // 打开链接
        consumer.connect();
        Thread.sleep(10000000);
    }
    public static void main(String[] args) throws Exception {
        produce1();
    }

//    @Test
    public  static void produce1() throws Exception {
        DmqBuilder builder =  DMQFactory.newBuilder("127.0.0.1",9208);

        // 创建一个消费者
        DmqObject  producer =  builder.buildProducer("phone","hxBA7bZlcJKE72kZ1adLmdnjhGx0Q+PgM3joSeLxcTM=");
        producer.receive( response -> {
            System.out.println(response.getData().getMessage());
        });


        // 打开链接
        producer.connect();
        Scanner scanner = new Scanner(System.in);
        FileInputStream in = new FileInputStream("D:\\高东昇的宝库\\source\\source\\www.csdn.net.sql");
        BufferedReader in1 = new BufferedReader(new InputStreamReader(in));
        int i = 0;
        while(true){

            String a= in1.readLine();
            i++;
            if(i ==20000){
                Thread.sleep(10000000);
                break;
            }
            producer.send(a);
//            producer.end();
            System.out.printf("%d  %s\n",i,a);
        }
//        FileInputStream in = new FileInputStream(ClassLoader.getSystemResource("www.csdn.net.sql").getPath());
//        BufferedReader in1 = new BufferedReader(new InputStreamReader(in));
//        while (true){
//
//            String a= in1.readLine();
//            i++;
//            if(i ==22){
//                break;
//            }
//            producer.send(a);
////            producer.end();
//            System.out.printf("%d  %s\n",i,a);
//        }

    }
//    @Test
    public static   void produce() throws Exception {
        DmqBuilder builder =  DMQFactory.newBuilder("127.0.0.1",9201);

        // 创建一个消费者
        DmqObject  producer =  builder.buildProducer("default","hxBA7bZlcJKE72kZ1adLmdnjhGx0Q+PgM3joSeLxcTM=");
        producer.receive( response -> {
            System.out.println(response.getData().getMessage());
        });


        // 打开链接
        producer.connect();
        Scanner scanner = new Scanner(System.in);
        while(true){
            System.out.println("========");
            String msg = scanner.next();
            producer.send(msg);
        }
    }
}
