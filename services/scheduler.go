package services

import (
	"math"
	"sync"
	"time"

	"github.com/scheduler/models"
)

type Scheduler struct {
	JobsList   []models.Job
	runningJob struct {
		ml     sync.Mutex
		jobLoc *models.Job
	}
	isSomeJobPending bool
}

var JobScheduler = Scheduler{
	JobsList:         []models.Job{},
	isSomeJobPending: true,
}

func updateStatusOfCurrentRunningProcess(status string) {
	JobScheduler.runningJob.ml.Lock()
	defer JobScheduler.runningJob.ml.Unlock()
	JobScheduler.runningJob.jobLoc.Status = status
}

func countDownJobTime() {
	JobScheduler.runningJob.ml.Lock()
	defer JobScheduler.runningJob.ml.Unlock()
	JobScheduler.runningJob.jobLoc.RemainingTime -= 1
}

func updateRunnigJob(currentJobToRun *models.Job) {
	JobScheduler.runningJob.ml.Lock()
	defer JobScheduler.runningJob.ml.Unlock()
	JobScheduler.runningJob.jobLoc = currentJobToRun
}

func RescheduleJobs() {
	shortestRemJob := int(math.Inf(1))
	indexOfShortestJob := -1
	for index, job := range JobScheduler.JobsList {
		if job.RemainingTime != 0 && job.RemainingTime < shortestRemJob {
			shortestRemJob = job.RemainingTime
			indexOfShortestJob = index
		}
	}

	if indexOfShortestJob >= 0 && JobScheduler.JobsList[indexOfShortestJob].RemainingTime != 0 {
		updateRunnigJob(&JobScheduler.JobsList[indexOfShortestJob])
		updateStatusOfCurrentRunningProcess("running")
		JobScheduler.isSomeJobPending = true
	} else {
		updateRunnigJob(nil)
		JobScheduler.isSomeJobPending = false
	}
}

func JobRunner() {
	RescheduleJobs()

	for {
		if JobScheduler.runningJob.jobLoc != nil && JobScheduler.runningJob.jobLoc.Status == "running" {
			for JobScheduler.runningJob.jobLoc.RemainingTime != 0 {
				time.Sleep(1 * time.Second)
				countDownJobTime()
			}
			updateStatusOfCurrentRunningProcess("completed")
		}

		if JobScheduler.isSomeJobPending {
			RescheduleJobs()
			go SendToAllClients([]byte("update_jobs_list"))
		}

	}

}

func GetJobsList() []models.Job {
	return JobScheduler.JobsList
}

func GetNumberOfJobs() int {
	return len(JobScheduler.JobsList)
}

func AddJob(newJob models.Job) {
	if JobScheduler.runningJob.jobLoc != nil && JobScheduler.runningJob.jobLoc.Status == "running" {
		updateStatusOfCurrentRunningProcess("pending")
	}
	JobScheduler.JobsList = append(JobScheduler.JobsList, newJob)
	RescheduleJobs()
}
