package com.example.grpcmicroservicetest

import com.google.protobuf.util.JsonFormat
import io.grpc.ManagedChannelBuilder
import io.grpc.ServerBuilder
import io.grpc.stub.StreamObserver
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.scheduling.annotation.EnableScheduling
import org.springframework.scheduling.annotation.Scheduled
import org.springframework.stereotype.Component
import java.lang.Exception
import java.time.Instant

@SpringBootApplication
@EnableScheduling
class GrpcmicroservicetestApplication

fun main(args: Array<String>) {
    runApplication<GrpcmicroservicetestApplication>(*args)
}

@Component
data class AppState(
    var counter: Long = 0L
)

class JavaGrpcServiceImpl(
    private val appState: AppState
) : JavaServiceGrpc.JavaServiceImplBase() {
    override fun increment(
        request: Mastermodel.EmptyMessage?,
        responseObserver: StreamObserver<Mastermodel.MasterMessage>?
    ) {
        appState.counter += 1
        println("[CALLED JavaGrpcServiceImpl] counter: ${appState.counter}")

        responseObserver?.onNext(
            Mastermodel.MasterMessage
                .newBuilder()
                .setPayload("[Java message] ${appState.counter}")
                .setTimestamp("${Instant.now()}")
                .build()
        )
        responseObserver?.onCompleted()
    }
}

@Component
class Runner(
    private val appState: AppState
) {
    init {
        try {
            ServerBuilder
                .forPort(50051)
                .addService(JavaGrpcServiceImpl(appState))
                .build()
                .start()
        } catch (e: Exception) {
            println("[GRPC SERVER ERR] $e")
        }
    }

    @Scheduled(fixedRate = 1)
    fun runTask() {
        println("[From Spring to Go service] ${Instant.now()}")

        try {
            val res = GoServiceGrpc.newBlockingStub(
                ManagedChannelBuilder
                    .forTarget("localhost:50052")
                    .usePlaintext()
                    .build()
            ).increment(
                Mastermodel
                    .EmptyMessage
                    .newBuilder()
                    .build()
            )

            println("Payload: ${JsonFormat.printer().print(res)}")
        } catch (e: Exception) {
            println("[error to go] $e")
        }
        println()

        Thread.sleep(5000)

        println("[From Spring to C# service] ${Instant.now()}")
        println("UNIMPLEMENTED")
        println()
        Thread.sleep(5000)
    }
}