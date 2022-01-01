declare var asu: Asu;

interface Asu {
    newLine(): Line;
    newTime(): Time;

    newTagB(): TagB;
    newTagPos(): TagPos;
    newTagPos(): TagPos;
    newTagMove(): TagMove;

    tagExists(text: string, tagType: string): boolean;
    replaceTag(text: string, tagType: string, newTag: string): string;
}

interface Error {
    Message: string;
}

interface Line {
    type: string;
    layer: number;
    start: Time;
    end: Time;
    style: string;
    actor: string;
    marginLeft: number;
    marginRight: number;
    marginVertical: number;
    effect: string;
    content: string;

    fromString(text: string): [Line | null, Error | null];
    toString(): string;
    length(): number;
    equals(): boolean;
}

interface Time {
    hours: number;
    minutes: number;
    seconds: number;

    fromSeconds(seconds: number): [Time | null, Error | null];
    fromArgs(hours: number, minutes: number, seconds: number): Time;
    fromString(text: string): [Time | null, Error | null];
    toString(): string;
    toSeconds(): number;
    add(t2: Time): Time;
    sub(t2: Time): Time;
    adjustOverplus(): void;
    equals(): boolean;
}

interface TagB {
    argument: number;

    name(): string;
    fromArgs(n: number): TagB;
    fromString(text: string): [TagB | null, Error | null];
    toString(): string;
}

interface TagPos {
    x: number;
    y: number;

    name(): string;
    fromArgs(x: number, y: number): TagPos;
    fromString(text: string): [TagPos | null, Error | null];
    toString(): string;
}
interface TagPos {
    x: number;
    y: number;

    name(): string;
    fromArgs(x: number, y: number): TagPos;
    fromString(text: string): [TagPos | null, Error | null];
    toString(): string;
}
interface TagMove {
    x1: number;
    y1: number;
    x2: number;
    y2: number;
    t1: number;
    t2: number;

    name(): string;
    fromArgs(x1: number, y1: number, x2: number, y2: number, t1: number, t2: number): TagMove;
    fromString(text: string): [TagMove | null, Error | null];
    toString(): string;
}
