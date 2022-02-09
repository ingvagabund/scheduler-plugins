package nodecrdresourcefit

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"

	crdclientset "sigs.k8s.io/scheduler-plugins/pkg/generated/clientset/versioned"
	crdinformers "sigs.k8s.io/scheduler-plugins/pkg/generated/informers/externalversions"
	v1alpha1schedulinginformers "sigs.k8s.io/scheduler-plugins/pkg/generated/informers/externalversions/scheduling/v1alpha1"
)

// NodeCRDResourceFit is a plugin which filters/scores nodes based on available
// resources on nodes provided through CRs
type NodeCRDResourceFit struct {
	handle      framework.Handle
	crdInformer v1alpha1schedulinginformers.ClusterScopedResorcesInformer
}

var _ = framework.ScorePlugin(&NodeCRDResourceFit{})
var _ = framework.FilterPlugin(&NodeCRDResourceFit{})

// NodeCRDResourceFitName is the name of the plugin used in the Registry and configurations.
const NodeCRDResourceFitName = "NodeCRDResourceFit"

// Name returns name of the plugin. It is used in logs, etc.
func (fit *NodeCRDResourceFit) Name() string {
	return NodeCRDResourceFitName
}

// NewNodeCRDResourceFit initializes a new plugin and returns it.
func NewNodeCRDResourceFit(fitArgs runtime.Object, handle framework.Handle) (framework.Plugin, error) {

	crdClient := crdclientset.NewForConfigOrDie(handle.KubeConfig())
	crdInformerFactory := crdinformers.NewSharedInformerFactory(crdClient, 0)
	crdInformer := crdInformerFactory.Scheduling().V1alpha1().ClusterScopedResorceses()

	ctx := context.TODO()
	crdInformerFactory.Start(ctx.Done())
	if !cache.WaitForCacheSync(ctx.Done(), crdInformer.Informer().HasSynced) {
		err := fmt.Errorf("WaitForCacheSync failed")
		klog.ErrorS(err, "Cannot sync caches")
		return nil, err
	}

	return &NodeCRDResourceFit{
		handle:      handle,
		crdInformer: crdInformer,
	}, nil
}

func (fit *NodeCRDResourceFit) Filter(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	// Return error if the node info is not properly populated
	if nodeInfo.Node() == nil {
		return framework.NewStatus(framework.Error, "node not found")
	}

	nodeCRDObject, err := fit.crdInformer.Lister().ClusterScopedResorceses("namespace").Get("nodename")
	if err != nil {
		return framework.NewStatus(framework.Error, "node CRD not found")
	}

	// implement filter logic

	return nil
}

func (fit *NodeCRDResourceFit) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	klog.V(5).InfoS("Scoring node", "nodeName", nodeName)

	// implement score logic

	return 0, nil
}

func (fit *NodeCRDResourceFit) ScoreExtensions() framework.ScoreExtensions {
	return nil
}
