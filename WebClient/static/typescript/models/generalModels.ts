import { Feeds } from '../rss/rss.models';

export class Settings {
    UnreadOnly: boolean;
    MarkSameRead: boolean;
    VkNewsEnabled: boolean;
    RssEnabled: boolean;
    ShowPreviewButton: boolean;
    ShowReadButton: boolean;
    ShowTabButton: boolean;
    VkLogin: string;
    VkPassword: string;
}

export class User {
    public Id: number;
    public Name: string;
    public Password: string;
    public VkLogin: string;
    public VkPassword: string;
    public VkNewsEnabled: string;
    public Settings: Settings;
}

export class RegistrationData {
    public User: User;
    public Message: string;
}

export class ModalData {
    Feed: Feeds;
    Settings: Settings;
}

export enum Sources {
    Rss = 1,
    Vk
}