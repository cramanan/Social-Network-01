"use client";

import { Event } from "@/types/group";
import { StrictOmit } from "@/utils/types";
import React, { ChangeEvent, FormEvent, useState } from "react";

type EventFields = StrictOmit<Event, "id" | "going">;

export default function NewEvent({ groupId }: { groupId: string }) {
    const defaultState = { groupId, title: "", description: "", date: "" };

    const [state, setState] = useState<EventFields>(defaultState);

    const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const response = await fetch("/api/create/event", {
            method: "POST",
            body: JSON.stringify(state),
        });
        if (response.ok) setState(defaultState);
    };

    const onChange = (key: keyof Event) => (e: ChangeEvent<HTMLInputElement>) =>
        setState({ ...state, [key]: e.target.value });

    return (
        <form onSubmit={onSubmit}>
            <label htmlFor="title">Title</label>
            <input
                type="text"
                name="title"
                id="title"
                value={state.title}
                onChange={onChange("title")}
            />
            <label htmlFor="description">Description</label>
            <input
                type="text"
                name="description"
                id="description"
                value={state.description}
                onChange={onChange("description")}
            />
            <label htmlFor="date">Date</label>
            <input
                type="datetime-local"
                name="date"
                id="date"
                value={state.date}
                onChange={onChange("date")}
            />
            <button type="submit">Create Event</button>
        </form>
    );
}
