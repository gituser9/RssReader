

export class ArticleData {
    Articles: Article[];
    Count: number;
}

export class Article {
    Id: number;
    Title: string;
    Body: string;
    Link: string;
    IsBookMark: boolean;
    IsRead: boolean;
}

export class Feed {
    ArticlesCount: number;
    ExistUnread: boolean;
    Feed: Feeds;
}

export class Feeds {
    Id: number;
    Url: string;
    Name: string;
    Articles: Article[];
}   

export class Tab {
    public title: string;
    public article: Article;
} 