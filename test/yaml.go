package test

import (
	"testing"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned/scheme"
	"k8s.io/apimachinery/pkg/runtime"
)

func mustParseYAML(t *testing.T, yaml string, i runtime.Object) {
	if _, _, err := scheme.Codecs.UniversalDeserializer().Decode([]byte(yaml), nil, i); err != nil {
		t.Fatalf("mustParseYAML (%s): %v", yaml, err)
	}
}

func mustParseTaskRun(t *testing.T, yaml string) *v1alpha1.TaskRun {
	var tr v1alpha1.TaskRun
	yaml = `apiVersion: tekton.dev/v1alpha1
kind: TaskRun
` + yaml
	mustParseYAML(t, yaml, &tr)
	return &tr
}

func mustParseTask(t *testing.T, yaml string) *v1alpha1.Task {
	var task v1alpha1.Task
	yaml = `apiVersion: tekton.dev/v1alpha1
kind: Task
` + yaml
	mustParseYAML(t, yaml, &task)
	return &task
}

func mustParsePipelineRun(t *testing.T, yaml string) *v1alpha1.PipelineRun {
	var pr v1alpha1.PipelineRun
	yaml = `apiVersion: tekton.dev/v1alpha1
kind: PipelineRun
` + yaml
	mustParseYAML(t, yaml, &pr)
	return &pr
}

func mustParsePipeline(t *testing.T, yaml string) *v1alpha1.Pipeline {
	var pipeline v1alpha1.Pipeline
	yaml = `apiVersion: tekton.dev/v1alpha1
kind: Pipeline
` + yaml
	mustParseYAML(t, yaml, &pipeline)
	return &pipeline
}
