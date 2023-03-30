package ernie_api

import (
	"context"
	"github.com/google/go-querystring/query"
	"net/http"
)

const (
	TaskPromptParagraph               = "PARAGRAPH"               //引导模型生成一段文章
	TaskPromptSend                    = "SENT"                    //引导模型生成一句话
	TaskPromptEntity                  = "ENTITY"                  //引导模型生成词组
	TaskPromptSummarization           = "Summarization"           //摘要
	TaskPromptMT                      = "MT"                      //翻译
	TaskPromptText2Annotation         = "Text2Annotation"         //抽取
	TaskPromptCorrection              = "Correction"              //纠错
	TaskPromptQAMRC                   = "QA_MRC"                  //阅读理解
	TaskPromptDialogue                = "Dialogue"                //对话
	TaskPromptQAClosedBook            = "QA_Closed_book"          //闭卷问答
	TaskPromptQAMultiChoice           = "QA_Multi_Choice"         //多选问答
	TaskPromptQuestionGeneration      = "QuestionGeneration"      //问题生成
	TaskPromptParaphrasing            = "Paraphrasing"            //复述
	TaskPromptNLI                     = "NLI"                     //文本蕴含识别
	TaskPromptSemanticMatching        = "SemanticMatching"        //匹配
	TaskPromptText2SQL                = "Text2SQL"                //文本描述转SQ
	TaskPromptTextClassification      = "TextClassification"      //文本分类
	TaskPromptSentimentClassification = "SentimentClassification" //情感分析
	TaskPromptZuoWen                  = "zuowen"                  //写作文
	TaskPromptAdText                  = "adtext"                  //写文案
	TaskPromptCouplet                 = "couplet"                 //对对联
	TaskPromptNovel                   = "novel"                   //写小说
	TaskPromptCloze                   = "cloze"                   //文本补全
	TaskPromptMisc                    = "Misc"                    //其它任务

)

type V3CustomizeRequest struct {
	Async             int     `json:"async" url:"async"`
	Text              string  `json:"text" url:"text"`
	MinDecLen         int     `json:"min_dec_len" url:"min_dec_len"`
	SeqLen            int     `json:"seq_len" url:"seq_len"`
	TopP              float32 `json:"topp" url:"topp"`
	PenaltyScore      float32 `json:"penalty_score,omitempty" url:"penalty_score,omitempty"`
	StopToken         string  `json:"stop_token,omitempty" url:"stop_token,omitempty"`
	TaskPrompt        string  `json:"task_prompt,omitempty" url:"task_prompt,omitempty"`
	TypeId            int     `json:"typeId" url:"typeId"`
	PenaltyText       string  `json:"penalty_text,omitempty" url:"penalty_text,omitempty"`
	ChoiceText        string  `json:"choice_text,omitempty" url:"choice_text,omitempty"`
	IsUnidirectional  int     `json:"is_unidirectional,omitempty" url:"is_unidirectional,omitempty"`
	MinDecPenaltyText string  `json:"min_dec_penalty_text,omitempty" url:"min_dec_penalty_text,omitempty"`
	LogitsBias        float32 `json:"logits_bias,omitempty" url:"logits_bias,omitempty"`
	MaskType          string  `json:"mask_type,omitempty" url:"mask_type,omitempty"`
}

type V3CustomizeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TaskID    int    `json:"taskId"`
		RequestID string `json:"requestId"`
	} `json:"data"`
}

func (c *Client) CreateV3Customize(ctx context.Context, request *V3CustomizeRequest) (response *V3CustomizeResponse, err error) {

	urlSuffix := "/rest/1.0/ernie/3.0.28/zeus"

	requestParams, err := query.Values(*request)
	if err != nil {
		return response, ErrV3CustomizeRequest
	}

	req, err := c.requestBuilder.build(ctx, http.MethodPost, c.fullURL(urlSuffix), requestParams)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
