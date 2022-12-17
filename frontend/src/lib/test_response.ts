import type { FeedData } from "./feed/types"

export const testResponse: FeedData = {
    id: 1,
    users: [
        {
            id: 1,
            firstName: "Max",
            lastName:"Mustermann",
            email: "abc",
            feed: [
                {
                    id: 123,
                    title: "Aufgabe 123",
                    notice: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, fugiat.",
                },
                {
                    id: 124,
                    title: "Aufgabe 124",
                    notice: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, fugiat.",
                },
                {
                    id: 123,
                    title: "Aufgabe 123",
                    notice: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, fugiat.",
                },
                {
                    id: 124,
                    title: "Aufgabe 124",
                    notice: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, fugiat.",
                },
                {
                    id: 123,
                    title: "Aufgabe 123",
                    notice: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, fugiat.",
                },
                {
                    id: 124,
                    title: "Aufgabe 124",
                    notice: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, fugiat.",
                }
            ]
        },
        {
            id: 2,
            firstName: "Marianne",
            lastName:"Musterfrau",
            email: "def",
            feed: [
                {
                    id: 125,
                    title: "Aufgabe 125",
                    notice: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, fugiat.",
                },
                {
                    id: 126,
                    title: "Aufgabe 126",
                    notice: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, fugiat.",
                }
            ]
        },
        {
            id: 3,
            firstName: "ABC",
            lastName:"DEEF",
            email: "ghi",
            feed: [
                {
                    id: 127,
                    title: "Aufgabe 127",
                    notice: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, fugiat.",
                }
            ]
        }
    ],
}