package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func InitTracer(serviceName, zipkinURL string) (func(context.Context) error, error) {
	// 1. Criar exporter do Zipkin
	exporter, err := zipkin.New(zipkinURL)
	if err != nil {
		return nil, err
	}

	// 2. Criar resource (identifica seu serviço)
	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, err
	}

	// 3. Criar TracerProvider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)

	// 4. Registrar globalmente
	otel.SetTracerProvider(tp)

	// 5. Retornar função de shutdown
	return tp.Shutdown, nil
}
