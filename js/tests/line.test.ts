import "../dist/asu";

test("lineContentAssignment", () => {
    const l = asu.newLine();
    const content = "some-text";
    l.content = content;
    expect(l.content).toBe(content);
});

test("newLineFromStringShouldNotFail", () => {
    const text = "Dialogue: 3,0:00:00.00,1:23:45.67,Default,Chitanda,1,2,3,Some fx,{\\b1\\fs32\\1c&H2F3066&}Some {\\1c&H0B0B26&}text";

    let expected = asu.newLine();
    expected.type = "Dialogue";
    expected.layer = 3;
    expected.start = asu.newTime().fromArgs(0, 0, 0);
    expected.end = asu.newTime().fromArgs(1, 23, 45.67);
    expected.style = "Default";
    expected.actor = "Chitanda";
    expected.marginLeft = 1;
    expected.marginRight = 2;
    expected.marginVertical = 3;
    expected.effect = "Some fx";
    expected.content = "{\\b1\\fs32\\1c&H2F3066&}Some {\\1c&H0B0B26&}text";

    const [actual, error] = asu.newLine().fromString(text);
    expect(error).toBeNull();
    expect(actual?.toString()).toStrictEqual(expected.toString());
});

test("newLineFromStringWithInvalidType", () => {
    const text = "Invalid: 3,0:00:00.00,1:23:45.67,Default,Chitanda,1,2,3,Some fx,{\\b1\\fs32\\1c&H2F3066&}Some {\\1c&H0B0B26&}text";

    let expectedError = asu.newError("invalid line type: 'Invalid'");

    const [actualLine, actualError] = asu.newLine().fromString(text);
    expect(actualError).not.toBe(null);
    expect(actualError?.message).toEqual(expectedError.message);
    expect(actualLine).toBeNull();
});
