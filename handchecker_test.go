package main




func TestGetAllPossibleHandsTest(t *testing.T) {
	in := struct{
		hand []card
		deck []card
	} {
		{[]card{card{faceTen, clubs}, card{face5, diamonds}},
		[]card{card{faceTen, clubs}, card{face5, diamonds}},
	}

	want := {}
	got :=

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAllPossibleHands(%q) == %q, want %q", in, got, want)
	}
}


[Test]
public void GetAllPossibleHandsTest()
{
	var actual =
		HandChecker.GetAllPossibleHands(CardUtil.GetCards("1H 2H 3H"), CardUtil.GetCards("1S 2S 3S"))
			.ToList();

	var expected = new List<IReadOnlyList<Card>>
	{
		CardUtil.GetCards("1H 2H 3H"),
		CardUtil.GetCards("1S 2H 3H"),
		CardUtil.GetCards("1H 1S 3H"),
		CardUtil.GetCards("1S 2S 3H"),
		CardUtil.GetCards("1H 2H 1S"),
		CardUtil.GetCards("1S 2H 2S"),
		CardUtil.GetCards("1H 1S 2S"),
		CardUtil.GetCards("1S 2S 3S"),
	};

	CollectionAssert.AreEqual(expected, actual);
}