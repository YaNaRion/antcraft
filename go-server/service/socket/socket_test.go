package socket

import (
	"main/infra"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SocketService", func() {
	var (
		socketService *SocketService
		db            infra.DBInterface
	)

	BeforeEach(func() {
		db = infra.MockDB{}
		socketService = NewSocketService(db)
	})

	Describe("GenerateRobotName", func() {
		It("should generate a unique robot name", func() {
			name1 := socketService.GenerateRobotName()
			socketService.Robots[name1] = "SOCKETID"
			name2 := socketService.GenerateRobotName()
			Expect(name1).NotTo(Equal(name2))
		})
	})

	Describe("FindRobotName", func() {
		Context("when the robot exists", func() {
			It("should return the robot name", func() {
				socketService.AddRobot("robot1", "socket1")
				name, err := socketService.FindRobotName("socket1")
				Expect(err).NotTo(HaveOccurred())
				Expect(name).To(Equal("robot1"))
			})
		})

		Context("when the robot does not exist", func() {
			It("should return an error", func() {
				_, err := socketService.FindRobotName("nonexistent")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("GetRobots", func() {
		It("should return the map of robots", func() {
			socketService.AddRobot("robot1", "socket1")
			socketService.AddRobot("robot2", "socket2")
			robots := socketService.GetRobots()
			Expect(robots).To(HaveLen(2))
			Expect(robots["robot1"]).To(Equal("socket1"))
			Expect(robots["robot2"]).To(Equal("socket2"))
		})
	})

	Describe("AddRobot", func() {
		It("should add a robot to the map", func() {
			socketService.AddRobot("robot1", "socket1")
			Expect(socketService.GetRobots()).To(HaveKeyWithValue("robot1", "socket1"))
		})
	})

	Describe("GetRobotNames", func() {
		It("should return a list of robot names", func() {
			socketService.AddRobot("robot1", "socket1")
			socketService.AddRobot("robot2", "socket2")
			names := socketService.GetRobotNames()
			Expect(names).To(ConsistOf("robot1", "robot2"))
		})
	})

	Describe("StartMission", func() {
		It("should start a mission", func() {
			socketService.StartMission(2, true)
			Expect(socketService.GetMissionStatus()).To(BeTrue())
			Expect(socketService.GetMission()).NotTo(BeNil())
		})
	})

	Describe("EndMission", func() {
		It("should end the mission and reset the state", func() {
			socketService.StartMission(2, true)
			socketService.EndMission()
			Expect(socketService.GetMissionStatus()).To(BeFalse())
			Expect(socketService.GetMission()).To(BeNil())
		})
	})

	Describe("GetMission", func() {
		It("should return the current mission", func() {
			socketService.StartMission(2, true)
			mission := socketService.GetMission()
			Expect(mission).NotTo(BeNil())
		})
	})

	Describe("GetMissionStatus", func() {
		It("should return the mission status", func() {
			socketService.StartMission(2, true)
			Expect(socketService.GetMissionStatus()).To(BeTrue())
			socketService.EndMission()
			Expect(socketService.GetMissionStatus()).To(BeFalse())
		})
	})
})
