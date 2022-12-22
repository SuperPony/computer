package behavioral.state;

public class Service {

    private ReplyStates state;

    public void reply(String userName) {
        Replier replier;
        switch (this.state) {
            case BUSY:
                replier = new BusyReplier();
                break;
            case LEISURE:
            default:
                replier = new LeisureReplier();
                break;
        }
        replier.reply(userName);
    }


    public ReplyStates getState() {
        return state;
    }

    public void setState(ReplyStates replyState) {
        this.state = replyState;
    }
}
