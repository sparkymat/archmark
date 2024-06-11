import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

dayjs.extend(utc);

export class Bookmark {
  public id: string;

  public createdAt: dayjs.Dayjs;

  public updatedAt: dayjs.Dayjs;

  public title: string;
  public transcript: string;
  public url: string;
  public description: string;
  public image?: string;
  public userId: string;

  constructor(json: any) {
    if (!json) {
      return;
    }

    this.id = json.id;
    this.createdAt = dayjs.utc(json.createdAt);
    this.updatedAt = dayjs.utc(json.updatedAt);

    this.title = json.title;
    this.url = json.url;
    this.description = json.description;
    this.transcript = json.transcript;

    if (json.image) {
      this.image = json.image;
    }
    this.userId = json.userId;
  }
}
