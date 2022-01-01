import "../dist/asu";

test("ParseTimeFromString", () => {
    const expected = "1:23:45.67"
    const [t, err] = asu.newTime().fromString(expected)
    const actual = t?.toString()
    expect(err).toBeNull()
    expect(actual).toEqual(expected)
});

test("ParseTimeFromArgs", () => {
    const expected = "1:23:45.67"
    const t = asu.newTime().fromArgs(1, 23, 45.67)
    const actual = t.toString()
    expect(actual).toEqual(expected)
});

test("Prints", () => {
    const t = asu.newTime()
    t.hours = 1
    t.minutes = 23
    t.seconds = 45.67
    const expected = "1:23:45.67"
    const actual = t.toString();
    expect(actual).toStrictEqual(expected);
});

test("PrintOverplusTime", () => {
    const t = asu.newTime()
    t.hours = 12
    t.minutes = 23
    t.seconds = 45.67
    const expected = "9:59:59.99"
    const actual = t.toString();
    expect(actual).toStrictEqual(expected);
});
