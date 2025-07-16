# Goods Customization Information

商品定制信息规范

## JSON

```json
{
    "raw_data": "[{\"name\":\"Joke\", \"age\":12}]",
    "surfaces": [
        {
            "name": null,
            "preview_image": {
                "label": null,
                "raw_url": "https://www.a.com/b.jpg",
                "url": "https://www.a.com/b.jpg",
                "valid": true,
                "error": null
            },
            "regions": []
        },
        {
            "name": null,
            "preview_image": {
                "label": null,
                "raw_url": "https://www.a.com/b.jpg",
                "url": "https://www.a.com/b.jpg",
                "valid": true,
                "error": null
            },
            "regions": [
                {
                    "name": "a",
                    "type": "text",
                    "texts": [
                        {
                            "label": "",
                            "value": "bbb"
                        }
                    ],
                    "images": [],
                    "valid": true,
                    "error": null
                }
            ]
        }
    ]
}
```