export type Group = {
    id: string;
    name: string;
    description: string;
    image: string;
    timestamp: string;
};

export type Event = {
    id: string;
    groupId: string;
    title: string;
    description: string;
    date: string;
    going: boolean;
};
