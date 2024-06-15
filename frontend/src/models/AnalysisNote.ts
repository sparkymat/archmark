import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';

dayjs.extend(utc);

export type Source = 'user' | 'service';

export class AnalysisNote {
  public id: string;

  public createdAt: dayjs.Dayjs;

  public updatedAt: dayjs.Dayjs;

  public content: string;

  public source: Source;

  public bookmarkId: string;

  constructor(json: any) {
    if (!json) {
      return;
    }

    this.id = json.id;
    this.createdAt = dayjs.utc(json.createdAt);
    this.updatedAt = dayjs.utc(json.updatedAt);

    this.content = json.content;
    this.source = json.source;
    this.bookmarkId = json.bookmarkId;
  }
}
