import "../dist/asu";

test("replaceTag", () => {
    const oldTag = "\\b1";
    const newTag = "\\pos(10,20)";
    const text = `{${oldTag}}Some text ${oldTag}{${oldTag}}Other text`;
    const expected = `{${newTag}}Some text ${oldTag}{${oldTag}}Other text`;
    const actual = asu.replaceTag(text, "b", newTag);
    expect(actual).toStrictEqual(expected);
});
