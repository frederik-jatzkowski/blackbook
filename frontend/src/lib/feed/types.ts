export type GroupData = {
    id: number,
    name: string
}

export type UserFeedData = {
    id: number,
    firstName: string,
    lastName: string,
    email: string,
    feed: TaskData[],
    commonGroups: GroupData[]
}

export type TaskData = {
    id: number,
    title: string,
    notice: string
}

export type FeedData = {
    users: UserFeedData[]
}