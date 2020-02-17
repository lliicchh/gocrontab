package worker

import (
	"context"
	"github.com/coreos/etcd/clientv3"
)

type JobLock struct {
	kv      clientv3.KV
	lease   clientv3.Lease
	jobName string
}

func InitLock(jobName string, kv clientv3.KV, lease clientv3.Lease) (jobLock *JobLock) {
	jobLock = &JobLock{
		kv:      kv,
		lease:   lease,
		jobName: jobName,
	}
	return jobLock
}

// 加锁
func (jobLock *JobLock) TryLock() (err error) {
	var (
		leaseGranResponse *clientv3.LeaseGrantResponse
	)
	if leaseGranResponse, err = jobLock.lease.Grant(context.TODO(), 4); err != nil {
		return
	}
}
