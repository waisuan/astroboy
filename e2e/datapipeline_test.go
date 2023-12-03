package e2e

//func TestDataPipeline(t *testing.T) {
//	a := assert.New(t)
//	os.Setenv("APP_ENV", "test")
//
//	deps := dependencies.Init()
//
//	user := model.User{
//		Username:    "esia",
//		Email:       "e-sia@outlook.com",
//		DateOfBirth: "12/11/1991",
//	}
//	uj, _ := json.Marshal(user)
//
//	m := model.Message{
//		T:       "user",
//		Key:     "esia",
//		Payload: string(uj),
//	}
//	mj, _ := json.Marshal(m)
//
//	err := deps.KafkaCli.ProduceMessage(kafka.Message{
//		Value: mj,
//	})
//	require.Nil(t, err)
//
//	found := false
//	for i := 0; i < 5; i++ {
//		req, err := http.NewRequest(http.MethodGet, "http://localhost:1324/api/users/esia", nil)
//		require.Nil(t, err)
//
//		res, err := http.DefaultClient.Do(req)
//		require.Nil(t, err)
//
//		if res != nil {
//			resBody, err := io.ReadAll(res.Body)
//			a.Nil(err)
//			a.Equal(http.StatusOK, res.StatusCode)
//			a.Equal(string(uj), strings.TrimSuffix(string(resBody), "\n"))
//
//			found = true
//			break
//		}
//
//		time.Sleep(time.Duration(i) * time.Second * 2)
//	}
//
//	if !found {
//		a.Fail("expected to receive a message but got nothing")
//	}
//}
